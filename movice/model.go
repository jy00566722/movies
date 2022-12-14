package movice

import (
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/qiniu/qmgo/field"
)

// Movie 结构体  百度资源
type Movice struct {
	// field.DefaultField `bson:",inline"`
	CreateTimeAt     time.Time `bson:"createTimeAt,omitempty"`
	UpdateTimeAt     time.Time `bson:"updateTimeAt,omitempty"`
	GroupId          int       `json:"group_id" form:"group_id" bson:"group_id"`
	TypeId           int       `json:"type_id" form:"type_id" bson:"type_id"`
	TypeId1          int       `json:"type_id_1" form:"type_id_1" bson:"type_id_1"`
	TypeName         string    `json:"type_name" form:"type_name" bson:"type_name"`
	VodActor         string    `json:"vod_actor" form:"vod_actor" bson:"vod_actor"`
	VodArea          string    `json:"vod_area" form:"vod_area" bson:"vod_area"`
	VodAuthor        string    `json:"vod_author" form:"vod_author" bson:"vod_author"`
	VodBehind        string    `json:"vod_behind" form:"vod_behind" bson:"vod_behind"`
	VodBlurb         string    `json:"vod_blurb" form:"vod_blurb" bson:"vod_blurb"`
	VodClass         string    `json:"vod_class" form:"vod_class" bson:"vod_class"`
	VodColor         string    `json:"vod_color" form:"vod_color" bson:"vod_color"`
	VodContent       string    `json:"vod_content" form:"vod_content" bson:"vod_content"`
	VodCopyright     int       `json:"vod_copyright" form:"vod_copyright" bson:"vod_copyright"`
	VodDirector      string    `json:"vod_director" form:"vod_director" bson:"vod_director"`
	VodDoubanId      int       `json:"vod_douban_id" form:"vod_douban_id" bson:"vod_douban_id"`
	VodDoubanScore   string    `json:"vod_douban_score" form:"vod_douban_score" bson:"vod_douban_score"`
	VodDown          int       `json:"vod_down" form:"vod_down" bson:"vod_down"`
	VodDownFrom      string    `json:"vod_down_from" form:"vod_down_from" bson:"vod_down_from"`
	VodDownNote      string    `json:"vod_down_note" form:"vod_down_note" bson:"vod_down_note"`
	VodDownServer    string    `json:"vod_down_server" form:"vod_down_server" bson:"vod_down_server"`
	VodDownUrl       string    `json:"vod_down_url" form:"vod_down_url" bson:"vod_down_url"`
	VodDuration      string    `json:"vod_duration" form:"vod_duration" bson:"vod_duration"`
	VodEn            string    `json:"vod_en" form:"vod_en" bson:"vod_en"`
	VodHits          int       `json:"vod_hits" form:"vod_hits" bson:"vod_hits"`
	VodHitsDay       int       `json:"vod_hits_day" form:"vod_hits_day" bson:"vod_hits_day"`
	VodHitsMonth     int       `json:"vod_hits_month" form:"vod_hits_month" bson:"vod_hits_month"`
	VodHitsWeek      int       `json:"vod_hits_week" form:"vod_hits_week" bson:"vod_hits_week"`
	VodId            int       `json:"vod_id" form:"vod_id" bson:"vod_id"`
	VodIsend         int       `json:"vod_isend" form:"vod_isend" bson:"vod_isend"`
	VodJumpurl       string    `json:"vod_jumpurl" form:"vod_jumpurl" bson:"vod_jumpurl"`
	VodLang          string    `json:"vod_lang" form:"vod_lang" bson:"vod_lang"`
	VodLetter        string    `json:"vod_letter" form:"vod_letter" bson:"vod_letter"`
	VodLevel         int       `json:"vod_level" form:"vod_level" bson:"vod_level"`
	VodLock          int       `json:"vod_lock" form:"vod_lock" bson:"vod_lock"`
	VodName          string    `json:"vod_name" form:"vod_name" bson:"vod_name"`
	VodPic           string    `json:"vod_pic" form:"vod_pic" bson:"vod_pic"`
	VodPicCcreenshot string    `json:"vod_pic_screenshot" form:"vod_pic_screenshot" bson:"vod_pic_screenshot"`
	VodPicSlide      string    `json:"vod_pic_slide" form:"vod_pic_slide" bson:"vod_pic_slide"`
	VodPicThumb      string    `json:"vod_pic_thumb" form:"vod_pic_thumb" bson:"vod_pic_thumb"`
	VodPlayFrom      string    `json:"vod_play_from" form:"vod_play_from" bson:"vod_play_from"`
	VodPlayNote      string    `json:"vod_play_note" form:"vod_play_note" bson:"vod_play_note"`
	VodPlayServer    string    `json:"vod_play_server" form:"vod_play_server" bson:"vod_play_server"`
	VodPlayUrl       string    `json:"vod_play_url" form:"vod_play_url" bson:"vod_play_url"`
	VodPlot          int       `json:"vod_plot" form:"vod_plot" bson:"vod_plot"`
	VodPlotDetail    string    `json:"vod_plot_detail" form:"vod_plot_detail" bson:"vod_plot_detail"`
	VodPlotName      string    `json:"vod_plot_name" form:"vod_plot_name" bson:"vod_plot_name"`
	Vod_points       int       `json:"vod_points" form:"vod_points" bson:"vod_points"`
	VodPointsDown    int       `json:"vod_points_down" form:"vod_points_down" bson:"vod_points_down"`
	VodPointsPlay    int       `json:"vod_points_play" form:"vod_points_play" bson:"vod_points_play"`
	VodPubdate       string    `json:"vod_pubdate" form:"vod_pubdate" bson:"vod_pubdate"`
	VodPwd           string    `json:"vod_pwd" form:"vod_pwd" bson:"vod_pwd"`
	VodPwdDown       string    `json:"vod_pwd_down" form:"vod_pwd_down" bson:"vod_pwd_down"`
	VodPwdDownUrl    string    `json:"vod_pwd_down_url" form:"vod_pwd_down_url" bson:"vod_pwd_down_url"`
	VodPwdPlay       string    `json:"vod_pwd_play" form:"vod_pwd_play" bson:"vod_pwd_play"`
	VodPwdPlayUrl    string    `json:"vod_pwd_play_url" form:"vod_pwd_play_url" bson:"vod_pwd_play_url"`
	VodPwdRrl        string    `json:"vod_pwd_url" form:"vod_pwd_url" bson:"vod_pwd_url"`
	VodRelArt        string    `json:"vod_rel_art" form:"vod_rel_art" bson:"vod_rel_art"`
	VodRelVod        string    `json:"vod_rel_vod" form:"vod_rel_vod" bson:"vod_rel_vod"`
	Vod_remarks      string    `json:"vod_remarks" form:"vod_remarks" bson:"vod_remarks"`
	VodReurl         string    `json:"vod_reurl" form:"vod_reurl" bson:"vod_reurl"`
	VodScore         string    `json:"vod_score" form:"vod_score" bson:"vod_score"`
	VodScoreSll      int       `json:"vod_score_all" form:"vod_score_all" bson:"vod_score_all"`
	VodScoreNum      int       `json:"vod_score_num" form:"vod_score_num" bson:"vod_score_num"`
	VodSerial        string    `json:"vod_serial" form:"vod_serial" bson:"vod_serial"`
	VodState         string    `json:"vod_state" form:"vod_state" bson:"vod_state"`
	VodStatus        int       `json:"vod_status" form:"vod_status" bson:"vod_status"`
	VodSub           string    `json:"vod_sub" form:"vod_sub" bson:"vod_sub"`
	VodTag           string    `json:"vod_tag" form:"vod_tag" bson:"vod_tag"`
	VodTime          string    `json:"vod_time" form:"vod_time" bson:"vod_time"`
	VodTimeAdd       int       `json:"vod_time_add" form:"vod_time_add" bson:"vod_time_add"`
	VodTimeHits      int       `json:"vod_time_hits" form:"vod_time_hits" bson:"vod_time_hits"`
	VodTimeMake      int       `json:"vod_time_make" form:"vod_time_make" bson:"vod_time_make"`
	VodTotal         int       `json:"vod_total" form:"vod_total" bson:"vod_total"`
	VodTpl           string    `json:"vod_tpl" form:"vod_tpl" bson:"vod_tpl"`
	VodTplDown       string    `json:"vod_tpl_down" form:"vod_tpl_down" bson:"vod_tpl_down"`
	VodTplPlay       string    `json:"vod_tpl_play" form:"vod_tpl_play" bson:"vod_tpl_play"`
	VodTrysee        int       `json:"vod_trysee" form:"vod_trysee" bson:"vod_trysee"`
	VodTv            string    `json:"vod_tv" form:"vod_tv" bson:"vod_tv"`
	VodUp            int       `json:"vod_up" form:"vod_up" bson:"vod_up"`
	VodVersion       string    `json:"vod_version" form:"vod_version" bson:"vod_version"`
	VodWeekday       string    `json:"vod_weekday" form:"vod_weekday" bson:"vod_weekday"`
	VodWriter        string    `json:"vod_writer" form:"vod_writer" bson:"vod_writer"`
	VodYear          string    `json:"vod_year" form:"vod_year" bson:"vod_year"`
	BzPic            string    `json:"bz_pic,omitempty" form:"bz_pic" bson:"bz_pic,omitempty"`
}

func (u *Movice) CustomFields() field.CustomFieldsBuilder {
	return field.NewCustom().SetCreateAt("CreateTimeAt").SetUpdateAt("UpdateTimeAt")
}

type MoviceResp struct {
	Code      int      `json:"code"`
	Limit     string   `json:"limit"`
	List      []Movice `json:"list"`
	Msg       string   `json:"msg"`
	Page      string   `json:"page"`
	Pagecount int      `json:"pagecount"`
	Total     int      `json:"total"`
}

type MoviceReq struct {
	Ac string `json:"ac"`
	Pg int    `json:"pg"`
	H  int    `json:"h"`
}

type LoadFileinfo struct {
	FileName string              `json:"fileName"`
	B2OutPut *s3.PutObjectOutput `json:"b2OutPut"`
	Size     int64               `json:"size,omitempty"`
}

//1080资源站的结构体
type Movice1080 struct {
	CreateTimeAt  time.Time `bson:"createTimeAt,omitempty"`
	UpdateTimeAt  time.Time `bson:"updateTimeAt,omitempty"`
	VodTime       string    `json:"vod_time" form:"vod_time" bson:"vod_time"`
	VodId         string    `json:"vod_id" form:"vod_id" bson:"vod_id"`
	VodName       string    `json:"vod_name" form:"vod_name" bson:"vod_name"`
	VodEnname     string    `json:"vod_enname" form:"vod_enname" bson:"vod_enname"`
	VodSub        string    `json:"vod_sub" form:"vod_sub" bson:"vod_sub"`
	VodLetter     string    `json:"vod_letter" form:"vod_letter" bson:"vod_letter"`
	VodColor      string    `json:"vod_color" form:"vod_color" bson:"vod_color"`
	VodTag        string    `json:"vod_tag" form:"vod_tag" bson:"vod_tag"`
	VodClass      string    `json:"vod_class" form:"vod_class" bson:"vod_class"`
	TypeId        string    `json:"type_id" form:"type_id" bson:"type_id"`
	TypeName      string    `json:"type_name" form:"type_name" bson:"type_name"`
	VodPic        string    `json:"vod_pic" form:"vod_pic" bson:"vod_pic"`
	VodLang       string    `json:"vod_lang" form:"vod_lang" bson:"vod_lang"`
	VodArea       string    `json:"vod_area" form:"vod_area" bson:"vod_area"`
	VodYear       string    `json:"vod_year" form:"vod_year" bson:"vod_year"`
	VodRemarks    string    `json:"vod_remarks" form:"vod_remarks" bson:"vod_remarks"`
	VodActor      string    `json:"vod_actor" form:"vod_actor" bson:"vod_actor"`
	VodDirector   string    `json:"vod_director" form:"vod_director" bson:"vod_director"`
	VodSerial     string    `json:"vod_serial" form:"vod_serial" bson:"vod_serial"`
	VodLock       string    `json:"vod_lock" form:"vod_lock" bson:"vod_lock"`
	VodLevel      string    `json:"vod_level" form:"vod_level" bson:"vod_level"`
	VodHits       string    `json:"vod_hits" form:"vod_hits" bson:"vod_hits"`
	VodHitsDay    string    `json:"vod_hits_day" form:"vod_hits_day" bson:"vod_hits_day"`
	VodHitsWeek   string    `json:"vod_hits_week" form:"vod_hits_week" bson:"vod_hits_week"`
	VodHitsMonth  string    `json:"vod_hits_month" form:"vod_hits_month" bson:"vod_hits_month"`
	VodDuration   string    `json:"vod_duration" form:"vod_duration" bson:"vod_duration"`
	VodUp         string    `json:"vod_up" form:"vod_up" bson:"vod_up"`
	VodDown       string    `json:"vod_down" form:"vod_down" bson:"vod_down"`
	VodScore      string    `json:"vod_score" form:"vod_score" bson:"vod_score"`
	VodScoreAll   string    `json:"vod_score_all" form:"vod_score_all" bson:"vod_score_all"`
	VodScoreNum   string    `json:"vod_score_num" form:"vod_score_num" bson:"vod_score_num"`
	VodPointsPlay string    `json:"vod_points_play" form:"vod_points_play" bson:"vod_points_play"`
	VodPointsDown string    `json:"vod_points_down" form:"vod_points_down" bson:"vod_points_down"`
	VodContent    string    `json:"vod_content" form:"vod_content" bson:"vod_content"`
	VodPlayFrom   string    `json:"vod_play_from" form:"vod_play_from" bson:"vod_play_from"`
	VodPlayNote   string    `json:"vod_play_note" form:"vod_play_note" bson:"vod_play_note"`
	VodPlayServer string    `json:"vod_play_server" form:"vod_play_server" bson:"vod_play_server"`
	VodPlayUrl    string    `json:"vod_play_url" form:"vod_play_url" bson:"vod_play_url"`
	VodDownFrom   string    `json:"vod_down_from" form:"vod_down_from" bson:"vod_down_from"`
	VodDownNote   string    `json:"vod_down_note" form:"vod_down_note" bson:"vod_down_note"`
	VodDownServer string    `json:"vod_down_server" form:"vod_down_server" bson:"vod_down_server"`
	VodDownUrl    string    `json:"vod_down_url" form:"vod_down_url" bson:"vod_down_url"`
	BzPic         string    `json:"bz_pic,omitempty" form:"bz_pic" bson:"bz_pic,omitempty"`
}

type Movice1080Resp struct {
	Code      int          `json:"code" form:"code" bson:"code"`
	Limit     int          `json:"limit" form:"limit" bson:"limit"`
	List      []Movice1080 `json:"list" form:"list" bson:"list"`
	Msg       string       `json:"msg" form:"msg" bson:"msg"`
	Page      int          `json:"page" form:"page" bson:"page"`
	Pagecount int          `json:"pagecount" form:"pagecount" bson:"pagecount"`
	Total     int          `json:"total" form:"total" bson:"total"`
}

//所有资源集合后的结构体，只选有用的字段
// type MoviceAll struct {
// 	CreateTimeAt time.Time `bson:"createTimeAt,omitempty"`
// 	UpdateTimeAt time.Time `bson:"updateTimeAt,omitempty"`
// 	VodId        string    `json:"vod_id" form:"vod_id" bson:"vod_id"`
// 	VodIds       struct {
// 		Bdzy   string `json:"bdzy,omitempty" form:"bdzy" bson:"bdzy,omitempty"`
// 		Gq1080 string `json:"gq1080,omitempty" form:"gq1080" bson:"gq1080,omitempty"`
// 	} `json:"vod_ids" form:"vod_ids" bson:"vod_ids"`
// 	VodName  string `json:"vod_name" form:"vod_name" bson:"vod_name"`
// 	TypeId   string `json:"type_id" form:"type_id" bson:"type_id"`
// 	TypeName struct {
// 		Bdzy   string `json:"bdzy,omitempty" form:"bdzy" bson:"bdzy,omitempty"`
// 		Qg1080 string `json:"gq1080,omitempty" form:"gq1080" bson:"gq1080,omitempty"`
// 	} `json:"type_name" form:"type_name" bson:"type_name"`
// 	VodPic struct {
// 		Bdzy   string `json:"bdzy,omitempty" form:"bdzy" bson:"bdzy,omitempty"`
// 		Qg1080 string `json:"gq1080,omitempty" form:"gq1080" bson:"gq1080,omitempty"`
// 	} `json:"vod_pic" form:"vod_pic" bson:"vod_pic"`
// 	VodArea     string `json:"vod_area" form:"vod_area" bson:"vod_area"`
// 	VodYear     string `json:"vod_year" form:"vod_year" bson:"vod_year"`
// 	VodRemarks  string `json:"vod_remarks" form:"vod_remarks" bson:"vod_remarks"`
// 	VodActor    string `json:"vod_actor" form:"vod_actor" bson:"vod_actor"`
// 	VodDirector string `json:"vod_director" form:"vod_director" bson:"vod_director"`
// 	VodContent  string `json:"vod_content" form:"vod_content" bson:"vod_content"`
// 	VodPlayUrl  struct {
// 		Bdzy   string `json:"bdzy,omitempty" form:"bdzy" bson:"bdzy,omitempty"`
// 		Gq1080 string `json:"gq1080,omitempty" form:"gq1080" bson:"gq1080,omitempty"`
// 	} `json:"vod_play_url" form:"vod_play_url" bson:"vod_play_url"`
// 	BzPic struct {
// 		Bdzy   string `json:"bdzy,omitempty" form:"bdzy" bson:"bdzy,omitempty"`
// 		Gq1080 string `json:"gq1080,omitempty" form:"gq1080" bson:"gq1080,omitempty"`
// 	} `json:"bz_pic,omitempty" form:"bz_pic" bson:"bz_pic,omitempty"`
// }

type MoviceMerge struct {
	VodName     string `json:"vod_name" form:"vod_name" bson:"vod_name"`
	VodIdBdzy   int    `json:"vod_id_bzdy" form:"vod_id_bzdy" bson:"vod_id_bzdy"`
	VodIdQg1080 string `json:"vod_id_gq1080" form:"vod_id_gq1080" bson:"vod_id_gq1080"`
}
