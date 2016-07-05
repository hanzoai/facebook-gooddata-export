package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/flosch/pongo2"
	"github.com/gocarina/gocsv"
	"github.com/hanzo-io/oauthful"
	"github.com/skratchdot/open-golang/open"
)

var redirectHTML = pongo2.Must(pongo2.FromFile("templates/redirect.html"))

type FacebookFlow struct {
}

func (f FacebookFlow) Decode(req *http.Request) (*oauthful.AuthorizationResponse, error) {
	errStr := req.FormValue("error")
	if errStr != "" {
		return nil, errors.New(errStr + "\n" + req.FormValue("error_reason") + "\n" + req.FormValue("error_description"))
	}

	res := &oauthful.AuthorizationResponse{
		State: req.FormValue("state"),
		Code:  req.FormValue("code"),
	}

	return res, nil
}

func (f FacebookFlow) Verify(res *oauthful.AuthorizationResponse) error {
	if res.State != Config.State {
		return errors.New("Authorization Error, State Does Not Match")
	}
	return nil
}

func (f FacebookFlow) AddParams(vals *url.Values) error {
	vals.Add("client_id", Config.ClientId)
	vals.Add("client_secret", Config.ClientSecret)
	vals.Add("redirect_uri", Config.RedirectUri)
	return nil
}

func FileExists(dir string) bool {
	if _, err := os.Stat(dir); err == nil {
		return true
	}
	return false
}

func MkDir(dir string) {
	if !FileExists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
}

func main() {
	toks, ok := getOAuthTokens()
	if !ok {
		fmt.Println("Getting new tokens.")
		toks, ok = newOAuthTokens()
		if !ok {
			fmt.Println("Could not get new tokens.")
			return
		}
	}

	csvExport(toks)
}

// Export all CSVs since last date (and update last one just in case)
func csvExport(toks *oauthful.AccessTokenResponse) {
	date := Date{Config.FirstDate}

	f, err := os.Open(Config.DataPath + "/date")
	if err == nil {
		err = Decode(f, &date)
		if err != nil {
			date.Date = Config.FirstDate
		}
	}

	MkDir(Config.ExportPath)
	now := time.Now()

	for now.After(date.Date) {
		fmt.Printf("Querying from Date %v\n", date.Date)

		ioutil.WriteFile(Config.DataPath+"/date", EncodeBytes(date), os.ModePerm)

		if err := writeFile(toks, Config.QueryByAudience, "audience", date.Date); err != nil {
			return
		}

		if err := writeFile(toks, Config.QueryByDevice, "device", date.Date); err != nil {
			return
		}

		if err := writeFile(toks, Config.QueryByCountry, "country", date.Date); err != nil {
			return
		}

		date.Date = date.Date.Add(time.Hour * 24)
	}

	mergeFiles("audience")
	mergeFiles("device")
	mergeFiles("country")
}

// Merge the CSV for a specific dataset
func mergeFiles(prefix string) error {
	exportPath := Config.ExportPath

	fileInfos, err := ioutil.ReadDir(exportPath + "/" + prefix)
	if err != nil {
		return err
	}

	fmt.Printf("Merging %v.csv\n", prefix)
	filename := exportPath + "/" + prefix + ".csv"

	_, _ = os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	fh, err := os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(fh)

	first := true
	for _, fileInfo := range fileInfos {
		fn := fileInfo.Name()

		if fn == ".DS_Store" {
			continue
		}

		f, _ := os.Open(exportPath + "/" + prefix + "/" + fn)
		defer f.Close()

		// fmt.Printf("Merge %v\n", fn)

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		if !first {
			scanner.Scan()
		} else {
			first = false
		}

		for scanner.Scan() {
			line := scanner.Text()
			// fmt.Printf("Line %v\n", line)
			fmt.Fprintln(w, line)
		}
	}
	w.Flush() // Don't forget to flush!

	return nil
}

// Write the CSV for a specific query for a specific date
func writeFile(toks *oauthful.AccessTokenResponse, q, prefix string, date time.Time) error {
	MkDir(Config.ExportPath + "/" + prefix)

	filename := Config.ExportPath + "/" + prefix + "/" + prefix + date.Format("_2006-01-02") + ".csv"
	// if prefix == "audience" || prefix == "device" {
	// 	if FileExists(filename) {
	// 		return nil
	// 	}
	// }

	insights := make([]FacebookInsight, 0)

	fmt.Printf("Using Query for %v\n", prefix)

	err := queryForDate(q, toks, date, &insights)
	if err != nil {
		fmt.Printf("Query Error for %v: %v\n", prefix, err)
		return err
	}

	csv, err := gocsv.MarshalString(insights)
	if err != nil {
		fmt.Printf("CSV Error for %v: %v\n", prefix, err)
		return err
	}
	ioutil.WriteFile(filename, []byte(csv), os.ModePerm)

	return nil
}

// Query Facebook Graph API for a specific date
func queryForDate(q string, toks *oauthful.AccessTokenResponse, date time.Time, insights *[]FacebookInsight) error {
	since := date.Add(time.Hour * -24)
	until := date

	insightsPart := make([]FacebookInsight, 0)

	url := fmt.Sprintf(q, Config.AppId, toks.AccessToken, since.Format("2006-01-02"), until.Format("2006-01-02"))

	asyncJob := FacebookAsyncJobResponse{}
	for {
		// Issue a POST request to start async job
		postUrl := Config.GraphUrl + "/" + url
		res, err := http.Post(postUrl, "application/json", nil)
		if err != nil {
			return err
		}

		err = Decode(res.Body, &asyncJob)
		if err != nil {
			return err
		}

		if asyncJob.Error.Message != "" {
			if strings.Contains(asyncJob.Error.Message, "#17") {
				fmt.Println("Out of API, waiting 10 minutes before retrying...")
				time.Sleep(time.Second * 600)
			}
			continue
		}

		break
	}

	// fmt.Printf("zzz %#v", postUrl)

	// Start polling for job completion
	id := asyncJob.ReportRunId
	asyncJobUrl := Config.GraphUrl + "/" + id + "?access_token=" + toks.AccessToken
	for {
		fmt.Printf("Waiting 5 seconds for Async Job %v to Complete...\n", id)
		time.Sleep(time.Second * 5)
		res, err := http.Get(asyncJobUrl)
		if err != nil {
			return err
		}

		// fmt.Printf("zzz %#v", asyncJobUrl)

		asyncJobStatus := FacebookAsyncJobStatus{}
		err = Decode(res.Body, &asyncJobStatus)
		if err != nil {
			return err
		}

		if asyncJobStatus.Error.Message != "" {
			if strings.Contains(asyncJobStatus.Error.Message, "#17") {
				fmt.Println("Out of API, waiting 10 minutes before retrying...")
				time.Sleep(time.Second * 600)
			}
			// return errors.New(asyncJobStatus.Error.Message)
			continue
		}

		fmt.Printf("Job Percent %v\n", asyncJobStatus.AsyncPercentCompletion)
		if asyncJobStatus.AsyncPercentCompletion == 100 {
			break
		}
	}

	// Wait a second so facebook can serve the report
	time.Sleep(time.Second)

	// Start querying data out
	wrapper := FacebookInsightWrapper{Paging: FacebookPaging{Next: Config.GraphUrl + "/" + id + "/insights?access_token=" + toks.AccessToken + "&limit=100"}}

	for wrapper.Paging.Next != "" {
		nextUrl := wrapper.Paging.Next
		// fmt.Printf("Url: %v\n", nextUrl)

		res, err := http.Get(nextUrl)
		if err != nil {
			return err
		}

		wrapper = FacebookInsightWrapper{}
		wrapper.Data = &insightsPart

		err = Decode(res.Body, &wrapper)
		if err != nil {
			return err
		}

		// Retry if out of api, otherwise its super annoying
		if wrapper.Error.Message != "" {
			if strings.Contains(wrapper.Error.Message, "#17") {
				wrapper.Paging.Next = nextUrl
				fmt.Println("Out of API, waiting 10 minutes before retrying...")
				time.Sleep(time.Second * 600)
				continue
			}
			return errors.New(wrapper.Error.Message)
		}

		for _, insight := range insightsPart {
			*insights = append((*insights), insight)
		}

		if wrapper.Paging.Next != "" {
			fmt.Println("Loading Next Page")
		} else {
			fmt.Println("All Pages Loaded")
		}
	}

	for i, insight := range *insights {
		(*insights)[i].ActionsStr = flattenJson(Encode(insight.Actions))
		(*insights)[i].WebsiteCtrStr = flattenJson(Encode(insight.WebsiteCtr))
	}

	return nil
}

// Flatten the JSON so it doesn't screw up the Good Data Parser
func flattenJson(str string) string {
	return strings.Replace(strings.Replace(strings.Replace(str, "\"", "", -1), " ", "", -1), "\n", "", -1)
}

// Load OAuth Tokens from File if possible
func getOAuthTokens() (*oauthful.AccessTokenResponse, bool) {
	f, err := os.Open(Config.TokensPath)
	if err != nil {
		return nil, false
	}

	fmt.Println("Loading Tokens")
	toks := &oauthful.AccessTokenResponse{}

	err = Decode(f, toks)
	if err != nil {
		return nil, false
	}

	fmt.Println("Decoding Tokens")
	return toks, true
}

// Issue an OAuth2.0 Authorization Request
func newOAuthTokens() (*oauthful.AccessTokenResponse, bool) {
	go server()

	qs := []string{
		"client_id=" + Config.ClientId,
		"state=" + Config.State,
		"scope=" + Config.Scope,
		"redirect_uri=" + Config.RedirectUri,
	}

	url := Config.AuthorizeUrl + "?" + strings.Join(qs, "&")

	open.Run(url)

	time.Sleep(time.Second * 2)

	return getOAuthTokens()
}

// OAuth2.0 Redirect Handler
func oauthRedirectHandler(w http.ResponseWriter, r *http.Request) {
	hc := &http.Client{}
	client := oauthful.New(hc, Config.TokenUrl, FacebookFlow{})

	pctx := pongo2.Context{
		"success": true,
	}

	if res, err := client.Handle(r); err != nil {
		pctx["success"] = false
		pctx["error"] = err.Error()
	} else {
		pctx["response"] = res
	}

	err := redirectHTML.ExecuteWriter(pctx, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	MkDir(Config.DataPath)
	ioutil.WriteFile(Config.TokensPath, EncodeBytes(pctx["response"]), os.ModePerm)
}

// Run the server that waits to execute the OAuth2.0 Redirect Handler
func server() {
	http.HandleFunc("/redirect", oauthRedirectHandler)
	http.ListenAndServe(":8080", nil)
}
