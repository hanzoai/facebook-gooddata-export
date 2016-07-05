package main

import "time"

type config struct {
	State string

	GraphUrl     string
	AuthorizeUrl string
	TokenUrl     string
	RedirectUri  string

	ClientId     string
	ClientSecret string
	Scope        string

	AppId string

	DataPath   string
	TokensPath string
	ExportPath string

	FirstDate time.Time

	QueryByAudience string
	QueryByDevice   string
	QueryByCountry  string
}

var Config = config{
	State: "12345678",

	GraphUrl:     "https://graph.facebook.com/v2.6",
	AuthorizeUrl: "https://www.facebook.com/dialog/oauth",
	TokenUrl:     "https://graph.facebook.com/v2.3/oauth/access_token",
	RedirectUri:  "http://localhost:8080/redirect",

	ClientId:     "244032362636714",
	ClientSecret: "b680c55a6ef62e603feaf1f33c0e15f1",
	Scope:        "read_insights,ads_read",

	AppId: "416885995169392",

	DataPath:   "./data",
	TokensPath: "./data/tokens",
	ExportPath: "./exports",

	FirstDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),

	QueryByAudience: "act_%v/insights?access_token=%v&time_range={'since':'%v','until':'%v'}&level=ad&breakdowns=age,gender&fields=account_id,account_name,action_values,actions,ad_id,ad_name,adset_id,adset_name,app_store_clicks,buying_type,call_to_action_clicks,campaign_id,campaign_name,canvas_avg_view_percent,canvas_avg_view_time,clicks,cost_per_10_sec_video_view,cost_per_action_type,cost_per_inline_link_click,cost_per_inline_post_engagement,cost_per_total_action,cost_per_unique_action_type,cost_per_unique_click,cost_per_unique_inline_link_click,cpc,cpm,cpp,ctr,deeplink_clicks,frequency,impressions,inline_link_click_ctr,inline_link_clicks,inline_post_engagement,newsfeed_avg_position,newsfeed_clicks,newsfeed_impressions,objective,place_page_name,reach,relevance_score,social_clicks,social_impressions,social_reach,social_spend,spend,total_action_value,total_actions,total_unique_actions,unique_actions,unique_ctr,unique_impressions,unique_inline_link_click_ctr,unique_inline_link_clicks,unique_link_clicks_ctr,unique_social_clicks,unique_social_impressions,video_10_sec_watched_actions,video_15_sec_watched_actions,video_30_sec_watched_actions,video_avg_pct_watched_actions,video_avg_sec_watched_actions,video_complete_watched_actions,video_p100_watched_actions,video_p25_watched_actions,video_p50_watched_actions,video_p75_watched_actions,video_p95_watched_actions,website_clicks,website_ctr",
	QueryByDevice:   "act_%v/insights?access_token=%v&time_range={'since':'%v','until':'%v'}&level=ad&breakdowns=placement,impression_device&fields=account_id,account_name,action_values,actions,ad_id,ad_name,adset_id,adset_name,app_store_clicks,buying_type,call_to_action_clicks,campaign_id,campaign_name,canvas_avg_view_percent,canvas_avg_view_time,clicks,cost_per_10_sec_video_view,cost_per_action_type,cost_per_inline_link_click,cost_per_inline_post_engagement,cost_per_total_action,cost_per_unique_action_type,cost_per_unique_click,cost_per_unique_inline_link_click,cpc,cpm,cpp,ctr,deeplink_clicks,frequency,impressions,inline_link_click_ctr,inline_link_clicks,inline_post_engagement,newsfeed_avg_position,newsfeed_clicks,newsfeed_impressions,objective,place_page_name,reach,relevance_score,social_clicks,social_impressions,social_reach,social_spend,spend,total_action_value,total_actions,total_unique_actions,unique_actions,unique_ctr,unique_impressions,unique_inline_link_click_ctr,unique_inline_link_clicks,unique_link_clicks_ctr,unique_social_clicks,unique_social_impressions,video_10_sec_watched_actions,video_15_sec_watched_actions,video_30_sec_watched_actions,video_avg_pct_watched_actions,video_avg_sec_watched_actions,video_complete_watched_actions,video_p100_watched_actions,video_p25_watched_actions,video_p50_watched_actions,video_p75_watched_actions,video_p95_watched_actions,website_clicks,website_ctr",

	// QueryByCountry:  "act_%v/insights?access_token=%v&time_range={'since':'%v','until':'%v'}act_416885995169392/insights?breakdowns=country&time_range={'since':'2016-06-29','until':'2016-06-30'}&level=ad&breakdowns=country&fields=account_id,account_name,action_values,actions,ad_id,ad_name,adset_id,adset_name,app_store_clicks,buying_type,call_to_action_clicks,campaign_id,campaign_name,canvas_avg_view_percent,canvas_avg_view_time,clicks,cost_per_10_sec_video_view,cost_per_action_type,cost_per_inline_link_click,cost_per_inline_post_engagement,cost_per_total_action,cost_per_unique_action_type,cost_per_unique_click,cost_per_unique_inline_link_click,cpc,cpm,cpp,ctr,deeplink_clicks,frequency,impressions,inline_link_click_ctr,inline_link_clicks,inline_post_engagement,newsfeed_avg_position,newsfeed_clicks,newsfeed_impressions,objective,place_page_name,reach,relevance_score,social_clicks,social_impressions,social_reach,social_spend,spend,total_action_value,total_actions,total_unique_actions,unique_actions,unique_ctr,unique_impressions,unique_inline_link_click_ctr,unique_inline_link_clicks,unique_link_clicks_ctr,unique_social_clicks,unique_social_impressions,website_clicks,website_ctr&limit=20",

	QueryByCountry: "act_%v/insights?access_token=%v&time_range={'since':'%v','until':'%v'}&level=ad&breakdowns=country&fields=account_id,account_name,action_values,actions,ad_id,ad_name,adset_id,adset_name,app_store_clicks,buying_type,call_to_action_clicks,campaign_id,campaign_name,canvas_avg_view_percent,canvas_avg_view_time,clicks,cost_per_10_sec_video_view,cost_per_action_type,cost_per_inline_link_click,cost_per_inline_post_engagement,cost_per_total_action,cost_per_unique_action_type,cost_per_unique_click,cost_per_unique_inline_link_click,cpc,cpm,cpp,ctr,deeplink_clicks,frequency,impressions,inline_link_click_ctr,inline_link_clicks,inline_post_engagement,newsfeed_avg_position,newsfeed_clicks,newsfeed_impressions,objective,place_page_name,reach,relevance_score,social_clicks,social_impressions,social_reach,social_spend,spend,total_action_value,total_actions,total_unique_actions,unique_actions,unique_ctr,unique_impressions,unique_inline_link_click_ctr,unique_inline_link_clicks,unique_link_clicks_ctr,unique_social_clicks,unique_social_impressions,video_10_sec_watched_actions,video_15_sec_watched_actions,video_30_sec_watched_actions,video_avg_pct_watched_actions,video_avg_sec_watched_actions,video_complete_watched_actions,video_p100_watched_actions,video_p25_watched_actions,video_p50_watched_actions,video_p75_watched_actions,video_p95_watched_actions,website_clicks,website_ctr",
}
