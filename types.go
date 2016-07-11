package main

import "time"

type AdsActionStats struct {
	// Click1d  float64 `json:"1d_click"`
	// View1d   float64 `json:"1d_view"`
	// Click28d float64 `json:"28d_click"`
	// View28d  float64 `json:"28d_view"`
	// Click7d  float64 `json:"7d_click"`
	// View7d   float64 `json:"7d_view"`

	// ActionCarouselCardId   string  `json:"action_carousel_card_id"`
	// ActionCarouselCardName string  `json:"action_carousel_card_name"`
	// ActionDestination      string  `json:"action_destination"`
	// ActionDevice           string  `json:"action_device"`
	// ActionTargetId         string  `json:"action_target_id"`
	ActionType string `json:"action_type"`
	// ActionVideoType        string  `json:"action_video_type"`
	Value float64 `json:"value"`
}

type FacebookAsyncJobResponse struct {
	ReportRunId string        `json:"report_run_id"`
	Error       FacebookError `json:"error"`
}

type FacebookAsyncJobStatus struct {
	Id                     string        `json:"id"`
	AsyncPercentCompletion int64         `json:"async_percent_completion"`
	Error                  FacebookError `json:"error"`
}

type Date struct {
	Date time.Time `json:"date"`
}

type FacebookError struct {
	Message     string `json:"message"`
	Type        string `json:"type"`
	Code        int64  `json:"code"`
	IsTransient bool   `json:"is_transient"`
	FBTraceId   string `json:"fbtrace_id"`
}

type FacebookPaging struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type FacebookInsightWrapper struct {
	Data   *[]FacebookInsight `json:"data"`
	Error  FacebookError      `json:"error"`
	Paging FacebookPaging     `json:"paging"`
}

type FacebookInsight struct {
	AccountId                                 string `json:"account_id,omitempty" csv:"account_id"`
	AccountName                               string `json:"account_name,omitempty" csv:"account_name"`
	AdId                                      string `json:"ad_id,omitempty" csv:"ad_id"`
	AdName                                    string `json:"ad_name,omitempty" csv:"ad_name"`
	AdsetId                                   string `json:"adset_id,omitempty" csv:"adset_id"`
	AdsetName                                 string `json:"adset_name,omitempty" csv:"adset_name"`
	Age                                       string `json:"age,omitempty" csv:"age"`
	BuyingType                                string `json:"buying_type,omitempty" csv:"buying_type"`
	CampaignId                                string `json:"campaign_id,omitempty" csv:"campaign_id"`
	CampaignName                              string `json:"campaign_name,omitempty" csv:"campaign_name"`
	Country                                   string `json:"country,omitempty" csv:"country"`
	FrequencyValue                            string `json:"frequency_value,omitempty" csv:"frequency_value"`
	Gender                                    string `json:"gender,omitempty" csv:"gender"`
	HourlyStatsAggregatedByAdvertiserTimeZone string `json:"hourly_stats_aggregated_by_advertiser_time_zone,omitempty" csv:"hourly_stats_aggregated_by_advertiser_time_zone"`
	HourlyStatsAggregatedByAudienceTimeZone   string `json:"hourly_stats_aggregated_by_audience_time_zone,omitempty" csv:"hourly_stats_aggregated_by_audience_time_zone"`
	ImpressionDevice                          string `json:"impression_device,omitempty" csv:"impression_device"`
	Impressions                               string `json:"impressions,omitempty" csv:"impressions"`
	Objective                                 string `json:"objective,omitempty" csv:"objective"`
	PlacePageId                               string `json:"place_page_id,omitempty" csv:"place_page_id"`
	PlacePageName                             string `json:"place_page_name,omitempty" csv:"place_page_name"`
	Placement                                 string `json:"placement,omitempty" csv:"placement"`
	ProductId                                 string `json:"product_id,omitempty" csv:"product_id"`
	Region                                    string `json:"region,omitempty" csv:"region"`

	DateStart string `json:"date_start,omitempty" csv:"date_start"`
	DateStop  string `json:"date_stop,omitempty" csv:"date_stop"`

	AppStoreClicks               int64   `json:"app_store_clicks,omitempty" csv:"app_store_clicks"`
	CallToActionClicks           int64   `json:"call_to_action_clicks,omitempty" csv:"call_to_action_clicks"`
	CanvasAvgViewPercent         float64 `json:"canvas_avg_view_percent,omitempty" csv:"canvas_avg_view_percent"`
	CanvasAvgViewTime            float64 `json:"canvas_avg_view_time,omitempty" csv:"canvas_avg_view_time"`
	Clicks                       int64   `json:"clicks,omitempty" csv:"clicks"`
	CostPerInlineLinkClick       float64 `json:"cost_per_inline_link_click,omitempty" csv:"cost_per_inline_link_click"`
	CostPerInlinePostEngagement  float64 `json:"cost_per_inline_post_engagement,omitempty" csv:"cost_per_inline_post_engagement"`
	CostPerTotalAction           float64 `json:"cost_per_total_action,omitempty" csv:"cost_per_total_action"`
	CostPerUniqueClick           float64 `json:"cost_per_unique_click,omitempty" csv:"cost_per_unique_click"`
	CostPerUniqueInlineLinkClick float64 `json:"cost_per_unique_inline_link_click,omitempty" csv:"cost_per_unique_inline_link_click"`
	Cpc                          float64 `json:"cpc,omitempty" csv:"cpc"`
	Cpm                          float64 `json:"cpm,omitempty" csv:"cpm"`
	Cpp                          float64 `json:"cpp,omitempty" csv:"cpp"`
	Ctr                          float64 `json:"ctr,omitempty" csv:"ctr"`
	DeeplinkClicks               int64   `json:"deeplink_clicks,omitempty" csv:"deeplink_clicks"`
	Frequency                    float64 `json:"frequency,omitempty" csv:"frequency"`
	InlineLinkClickCtr           float64 `json:"inline_link_click_ctr,omitempty" csv:"inline_link_click_ctr"`
	InlineLinkClicks             int64   `json:"inline_link_clicks,omitempty" csv:"inline_link_clicks"`
	InlinePostEngagement         int64   `json:"inline_post_engagement,omitempty" csv:"inline_post_engagement"`
	NewsfeedAvgPosition          float64 `json:"newsfeed_avg_position,omitempty" csv:"newsfeed_avg_position"`
	NewsfeedClicks               int64   `json:"newsfeed_clicks,omitempty" csv:"newsfeed_clicks"`
	NewsfeedImpressions          int64   `json:"newsfeed_impressions,omitempty" csv:"newsfeed_impressions"`
	Reach                        int64   `json:"reach,omitempty" csv:"reach"`
	SocialClicks                 int64   `json:"social_clicks,omitempty" csv:"social_clicks"`
	SocialImpressions            int64   `json:"social_impressions,omitempty" csv:"social_impressions"`
	SocialReach                  int64   `json:"social_reach,omitempty" csv:"social_reach"`
	SocialSpend                  float64 `json:"social_spend,omitempty" csv:"social_spend"`
	Spend                        float64 `json:"spend,omitempty" csv:"spend"`
	TotalActionValue             float64 `json:"total_action_value,omitempty" csv:"total_action_value"`
	TotalActions                 int64   `json:"total_actions,omitempty" csv:"total_actions"`
	TotalUniqueActions           int64   `json:"total_unique_actions,omitempty" csv:"total_unique_actions"`
	UniqueClicks                 int64   `json:"unique_clicks,omitempty" csv:"unique_clicks"`
	UniqueCtr                    float64 `json:"unique_ctr,omitempty" csv:"unique_ctr"`
	UniqueImpressions            int64   `json:"unique_impressions,omitempty" csv:"unique_impressions"`
	UniqueInlineLinkClickCtr     float64 `json:"unique_inline_link_click_ctr,omitempty" csv:"unique_inline_link_click_ctr"`
	UniqueInlineLinkClicks       int64   `json:"unique_inline_link_clicks,omitempty" csv:"unique_inline_link_clicks"`
	UniqueLinkClicksCtr          float64 `json:"unique_link_clicks_ctr,omitempty" csv:"unique_link_clicks_ctr"`
	UniqueSocialClicks           int64   `json:"unique_social_clicks,omitempty" csv:"unique_social_clicks"`
	UniqueSocialImpressions      int64   `json:"unique_social_impressions,omitempty" csv:"unique_social_impressions"`
	WebsiteClicks                int64   `json:"website_clicks,omitempty" csv:"website_clicks"`

	// ActionValues                []AdsActionStats `json:"action_values,omitempty" csv:"action_values"`
	Actions    []AdsActionStats `json:"actions,omitempty" csv:"-"`
	ActionsStr string           `csv:"actions"`
	// CostPer10SecVideoView       []AdsActionStats `json:"cost_per_10_sec_video_view,omitempty" csv:"cost_per_10_sec_video_view"`
	// CostPerActionType           []AdsActionStats `json:"cost_per_action_type,omitempty" csv:"cost_per_action_type"`
	// CostPerUniqueActionType     []AdsActionStats `json:"cost_per_unique_action_type,omitempty" csv:"cost_per_unique_action_type"`
	// UniqueActions               []AdsActionStats `json:"unique_actions,omitempty" csv:"unique_actions"`
	// Video10SecWatchedActions    []AdsActionStats `json:"video_10_sec_watched_actions,omitempty" csv:"video_10_sec_watched_actions"`
	// Video15SecWatchedActions    []AdsActionStats `json:"video_15_sec_watched_actions,omitempty" csv:"video_15_sec_watched_actions"`
	// Video30SecWatchedActions    []AdsActionStats `json:"video_30_sec_watched_actions,omitempty" csv:"video_30_sec_watched_actions"`
	// VideoAvgPctWatchedActions   []AdsActionStats `json:"video_avg_pct_watched_actions,omitempty" csv:"video_avg_pct_watched_actions"`
	// VideoAvgSecWatchedActions   []AdsActionStats `json:"video_avg_sec_watched_actions,omitempty" csv:"video_avg_sec_watched_actions"`
	// VideoCompleteWatchedActions []AdsActionStats `json:"video_complete_watched_actions,omitempty" csv:"video_complete_watched_actions"`
	// VideoP100WatchedActions     []AdsActionStats `json:"video_p100_watched_actions,omitempty" csv:"video_p100_watched_actions"`
	// VideoP25WatchedActions      []AdsActionStats `json:"video_p25_watched_actions,omitempty" csv:"video_p25_watched_actions"`
	// VideoP50WatchedActions      []AdsActionStats `json:"video_p50_watched_actions,omitempty" csv:"video_p50_watched_actions"`
	// VideoP75WatchedActions      []AdsActionStats `json:"video_p75_watched_actions,omitempty" csv:"video_p75_watched_actions"`
	// VideoP95WatchedActions      []AdsActionStats `json:"video_p95_watched_actions,omitempty" csv:"video_p95_watched_actions"`
	WebsiteCtr    []AdsActionStats `json:"website_ctr,omitempty" csv:"-"`
	WebsiteCtrStr string           `csv:"website_ctr"`

	CostPerActionType    []AdsActionStats `json:"cost_per_action_type" csv:"-"`
	CostPerActionTypeStr string           `csv:"cost_per_action_type"`
}
