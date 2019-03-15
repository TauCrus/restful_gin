package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"
)

// Copywriter 文案
type Copywriter struct {
}

// Banner 轮播图
type Banner struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	ImamgeURL    string `json:"image_url"`
	JumpURL      string `json:"jump_url"`
	Uptime       string `json:"uptime"`
	IsShow       int    `json:"is_show"`
	InReview     int    `json:"in_review"`
	ActivityName string `json:"activity_name"`
	PlaceID      int    `json:"place_id"`
	Place        string `json:"place"`
	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
}

// GetBanners 查询轮播图
func (c *Copywriter) GetBanners(isHp int, title, activityName, status string) (banners []Banner, err error) {
	banners = make([]Banner, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT b.id,
			b.title,
			b.image_url ,
			b.jump_url ,
			b.uptime,
			b.is_show ,
			b.in_review ,
			b.activity_name ,
			b.place_id,
			bp.place
		FROM gpxj_app.t_banner b
		LEFT JOIN gpxj_app.t_banner_place bp ON b.place_id = bp.id
		WHERE 1
	`)

	if 1 == isHp {
		querySQL = utils.SetSQLFormat(`{0} AND  b.place_id = 6`, querySQL)
	} else {
		querySQL = utils.SetSQLFormat(`{0} AND  b.place_id BETWEEN 1 AND 5`, querySQL)
	}

	if "" != title {
		querySQL = utils.SetSQLFormat(`{0} AND b.title like '%{1}%'`, querySQL, title)
	}

	if "" != activityName {
		querySQL = utils.SetSQLFormat(`{0} AND b.activity_name like '%{1}%'`, querySQL, activityName)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND b.is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND b.is_show = 0`, querySQL)
		} else if "3" == status {
			querySQL = utils.SetSQLFormat(`{0} AND b.in_review = 1`, querySQL)
		}
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY b.id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var banner Banner
		rows.Scan(&banner.ID, &banner.Title, &banner.ImamgeURL, &banner.JumpURL, &banner.Uptime,
			&banner.IsShow, &banner.InReview, &banner.ActivityName, &banner.PlaceID, &banner.Place)

		banners = append(banners, banner)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// BannerPlace 轮播图位置
type BannerPlace struct {
	ID    int    `json:"id"`
	Place string `json:"place"`
}

// GetBannerPlaces 查询轮播图位置
func (c *Copywriter) GetBannerPlaces() (bps []BannerPlace, err error) {
	bps = make([]BannerPlace, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id, place
		FROM t_banner_place 
		WHERE id IN (1,2,3,4,5)
	`)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var bp BannerPlace
		rows.Scan(&bp.ID, &bp.Place)

		bps = append(bps, bp)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// AddBanner 新增轮播图
func (c *Copywriter) AddBanner(banner Banner) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
	INSERT INTO gpxj_app.t_banner (
		column_id,title,image_url,jump_url,
		share_title,share_desc,share_icon_url,share_url,
		place_id,uptime,is_show,activity_name,in_review) 
	  VALUES
		(
		  '0','{0}','{1}','{2}',
		'{3}','{4}','{5}','{6}',
		  '{7}','{8}','{9}','{10}','{11}')
		  `, banner.Title, banner.ImamgeURL, banner.JumpURL,
		banner.ShareTitle, banner.ShareIconURL, banner.ShareURL, banner.ShareDesc,
		banner.PlaceID, banner.Uptime, banner.IsShow, banner.ActivityName, banner.InReview)

	log.Println("insertSQL:", insertSQL)

	rs, err := db.SQLDB.Exec(insertSQL)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}

	return
}

// ModifyBanner 修改轮播图
func (c *Copywriter) ModifyBanner(banner Banner) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
	UPDATE 
		gpxj_app.t_banner 
	SET
		title = '{0}',
		image_url = '{1}',
		jump_url = '{2}',
		share_title = '{3}',
		share_desc = '{4}',
		share_icon_url = '{5}',
		share_url = '{6}',
		place_id = '{7}',
		uptime = '{8}',
		is_show = '{9}',
		activity_name = '{10}',
		in_review = '{11}'
	WHERE id = '{12}' ;
	`, banner.Title, banner.ImamgeURL, banner.JumpURL,
		banner.ShareTitle, banner.ShareDesc, banner.ShareIconURL, banner.ShareURL,
		banner.PlaceID, banner.Uptime, banner.IsShow, banner.ActivityName, banner.InReview,
		banner.ID)

	log.Println("updateSQL:", updateSQL)

	stmt, err := db.SQLDB.Prepare(updateSQL)
	if nil != err {
		return
	}

	rs, err := stmt.Exec()
	if nil != err {
		return
	}

	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

// StartPage 启动页
type StartPage struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ImamgeURL string `json:"image_url"`
	JumpURL   string `json:"jump_url"`
	IsShow    int    `json:"is_show"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// GetStartPages 查询启动页
func (c *Copywriter) GetStartPages() (startpages []StartPage, err error) {
	startpages = make([]StartPage, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,title,image_url,jump_url,
			is_show,start_time,end_time 
		FROM gpxj_app.t_start_page 
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var startpage StartPage
		rows.Scan(&startpage.ID, &startpage.Title, &startpage.ImamgeURL, &startpage.JumpURL,
			&startpage.IsShow, &startpage.StartTime, &startpage.EndTime)

		startpages = append(startpages, startpage)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// SearchRecommend 搜索推荐
type SearchRecommend struct {
	ID            int    `json:"id"`
	StockName     string `json:"stock_name"`
	StockCode     string `json:"stock_code"`
	ActivityName  string `json:"activity_name"`
	JumpURL       string `json:"jump_url"`
	RecommendType string `json:"recommend_type"`
	IsShow        int    `json:"is_show"`
}

// GetSearchRecommends 查询搜索推荐
func (c *Copywriter) GetSearchRecommends() (srs []SearchRecommend, err error) {
	srs = make([]SearchRecommend, 0)

	querySQL := utils.SetSQLFormat(`
		 SELECT id,
		 	IFNULL(stock_name,""),IFNULL(stock_code,""),
 			IFNULL(activity_name,""),IFNULL(jump_url,""),
		 	recommend_type,is_show
 		FROM gpxj_app.t_search_recommend
 		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var sr SearchRecommend
		rows.Scan(&sr.ID, &sr.StockName, &sr.StockCode, &sr.ActivityName, &sr.JumpURL, &sr.RecommendType, &sr.IsShow)

		srs = append(srs, sr)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// MarketingLabel 行情营销标签
type MarketingLabel struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
	JumpURL string `json:"jump_url"`
	Place   string `json:"place"`
	IsShow  int    `json:"is_show"`
}

// GetMarketingLabels  查询行情营销标签
func (c *Copywriter) GetMarketingLabels() (labels []MarketingLabel, err error) {
	labels = make([]MarketingLabel, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,tag_name,jump_url,place,is_show
		FROM gpxj_app.t_marketing_label 
 		WHERE 1
		`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var label MarketingLabel
		rows.Scan(&label.ID, &label.TagName, &label.JumpURL, &label.Place, &label.IsShow)

		labels = append(labels, label)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}
