package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"

	"github.com/golang/glog"
)

// Copywriter 文案
type Copywriter struct {
}

// Banner 轮播图
type Banner struct {
	ID           int    `json:"id"`
	ColumnID     int    `json:"column_id"`
	ColumnName   string `json:"column_name"`
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
func (c *Copywriter) GetBanners(isHp int, keyword, placeID, columnID, status string) (banners []Banner, err error) {
	banners = make([]Banner, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT b.id,
			b.column_id,
			IFNULL(c.column_name,""),
			b.title,
			b.image_url ,
			b.jump_url ,
			DATE_FORMAT(uptime,'%Y-%m-%d %H:%i:%s'),
			b.is_show ,
			b.in_review ,
			b.activity_name ,
			b.place_id,
			bp.place,
			IFNULL(share_title,''),
			IFNULL(share_icon_url,''),
			IFNULL(share_url,''),
			IFNULL(share_desc,'')
		FROM gpxj_app.t_banner b
		LEFT JOIN gpxj_app.t_banner_place bp ON b.place_id = bp.id
		LEFT JOIN gpxj_app.t_column c ON b.column_id =  c.id
		WHERE 1
	`)

	if 1 == isHp {
		querySQL = utils.SetSQLFormat(`{0} AND  b.place_id = 6`, querySQL)
	} else {
		querySQL = utils.SetSQLFormat(`{0} AND  b.place_id BETWEEN 1 AND 5`, querySQL)
	}

	if "" != keyword {
		querySQL = utils.SetSQLFormat(`{0} AND (b.title like '%{1}%' OR b.activity_name like '%{1}%' )`, querySQL, keyword)
	}

	if "0" != placeID {
		querySQL = utils.SetSQLFormat(`{0} AND b.place_id = '{1}'`, querySQL, placeID)
	}

	if "0" != columnID {
		querySQL = utils.SetSQLFormat(`{0} AND b.column_id = '{1}'`, querySQL, columnID)
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

	glog.Info("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var banner Banner
		rows.Scan(&banner.ID, &banner.ColumnID, &banner.ColumnName,
			&banner.Title, &banner.ImamgeURL, &banner.JumpURL, &banner.Uptime,
			&banner.IsShow, &banner.InReview, &banner.ActivityName, &banner.PlaceID, &banner.Place,
			&banner.ShareTitle, &banner.ShareIconURL, &banner.ShareURL, &banner.ShareDesc)

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

	glog.Info("querySQL:", querySQL)

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
		  '{12}','{0}','{1}','{2}',
		'{3}','{4}','{5}','{6}',
		  '{7}','{8}','{9}','{10}','{11}')
		  `, banner.Title, banner.ImamgeURL, banner.JumpURL,
		banner.ShareTitle, banner.ShareIconURL, banner.ShareURL, banner.ShareDesc,
		banner.PlaceID, banner.Uptime, banner.IsShow, banner.ActivityName, banner.InReview,
		banner.ColumnID)

	glog.Info("insertSQL:", insertSQL)

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
			title = '{1}',
			image_url = '{2}',
			jump_url = '{3}',
			share_title = '{4}',
			share_desc = '{5}',
			share_icon_url = '{6}',
			share_url = '{7}',
			place_id = '{8}',
			uptime = '{9}',
			is_show = '{10}',
			activity_name = '{11}',
			in_review = '{12}',
			column_id = '{13}'
		WHERE id = '{0}' ;
		`, banner.ID, banner.Title, banner.ImamgeURL, banner.JumpURL,
		banner.ShareTitle, banner.ShareDesc, banner.ShareIconURL, banner.ShareURL,
		banner.PlaceID, banner.Uptime, banner.IsShow, banner.ActivityName, banner.InReview,
		banner.ColumnID)

	glog.Info("updateSQL:", updateSQL)

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

// DropBanner 删除轮播图
func (c *Copywriter) DropBanner(banner Banner) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_banner WHERE id ='{0}'`, banner.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
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
	ID           int    `json:"id"`
	Title        string `json:"title"`
	ImamgeURL    string `json:"image_url"`
	JumpURL      string `json:"jump_url"`
	IsShow       int    `json:"is_show"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
}

// GetStartPages 查询启动页
func (c *Copywriter) GetStartPages(title, status string) (startpages []StartPage, err error) {
	startpages = make([]StartPage, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,
			title,
			image_url,
			jump_url,
			is_show,
			IFNULL(share_title,''),
			IFNULL(share_icon_url,''),
			IFNULL(share_url,''),
			IFNULL(share_desc,''),
			DATE_FORMAT(start_time,'%Y-%m-%d %H:%i:%s'),
			DATE_FORMAT(end_time,'%Y-%m-%d %H:%i:%s')
		FROM gpxj_app.t_start_page 
		WHERE 1
	`)

	if "" != title {
		querySQL = utils.SetSQLFormat(`{0} AND title like '%{1}%'`, querySQL, title)
	}
	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 0`, querySQL)
		}
	}
	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	glog.Info("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var startpage StartPage
		rows.Scan(&startpage.ID, &startpage.Title, &startpage.ImamgeURL, &startpage.JumpURL, &startpage.IsShow,
			&startpage.ShareTitle, &startpage.ShareIconURL, &startpage.ShareURL, &startpage.ShareDesc,
			&startpage.StartTime, &startpage.EndTime)

		startpages = append(startpages, startpage)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// AddStartPage 新增启动页
func (c *Copywriter) AddStartPage(startpage StartPage) (id int64, err error) {
	insertSQL := utils.SetSQLFormat(`
		INSERT INTO gpxj_app.t_start_page (
			title,image_url,jump_url,is_show,
			share_title,share_icon_url,share_url,share_desc,
			start_time,end_time) 
		VALUES
			(
			'{0}','{1}','{2}','{3}',
			'{4}','{5}','{6}','{7}',
			'{8}','{9}');
		`, startpage.Title, startpage.ImamgeURL, startpage.JumpURL, startpage.IsShow,
		startpage.ShareTitle, startpage.ShareIconURL, startpage.ShareURL, startpage.ShareDesc,
		startpage.StartTime, startpage.EndTime)

	glog.Info("insertSQL:", insertSQL)

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

// ModifyStartPage 修改启动页
func (c *Copywriter) ModifyStartPage(startpage StartPage) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
		UPDATE 
			gpxj_app.t_start_page 
		SET
			title = '{1}',
			image_url = '{2}',
			jump_url = '{3}',
			is_show = '{4}',
			share_title = '{5}',
			share_desc = '{6}',
			share_icon_url = '{7}',
			share_url = '{8}',
			start_time = '{9}',
			end_time = '{10}'
		WHERE id = '{0}';  
		`, startpage.ID,
		startpage.Title, startpage.ImamgeURL, startpage.JumpURL, startpage.IsShow,
		startpage.ShareTitle, startpage.ShareIconURL, startpage.ShareURL, startpage.ShareDesc,
		startpage.StartTime, startpage.EndTime)

	glog.Info("updateSQL:", updateSQL)

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

// DropStartPage 删除启动页
func (c *Copywriter) DropStartPage(startpage StartPage) (id int64, err error) {
	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_start_page WHERE id ='{0}'`, startpage.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
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
	Sort          int    `json:"sort"`
	ShareTitle    string `json:"share_title"`
	ShareIconURL  string `json:"share_icon_url"`
	ShareURL      string `json:"share_url"`
	ShareDesc     string `json:"share_desc"`
}

// GetSearchRecommends 查询搜索推荐
func (c *Copywriter) GetSearchRecommends(recommendType, status string) (srs []SearchRecommend, err error) {
	srs = make([]SearchRecommend, 0)

	querySQL := utils.SetSQLFormat(`
		 SELECT id,
			 IFNULL(stock_name,""),
			 IFNULL(stock_code,""),
			 IFNULL(activity_name,""),
			 IFNULL(jump_url,""),
			 recommend_type,
			 is_show,
			 IFNULL(sort,0),
			 IFNULL(share_title,''),
			 IFNULL(share_icon_url,''),
			 IFNULL(share_url,''),
			 IFNULL(share_desc,'')
 		FROM gpxj_app.t_search_recommend
 		WHERE 1
	`)
	if "" != recommendType {
		querySQL = utils.SetSQLFormat(`{0} AND recommend_type = '{1}'`, querySQL, recommendType)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 0`, querySQL)
		}
	}
	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var sr SearchRecommend
		rows.Scan(&sr.ID, &sr.StockName, &sr.StockCode, &sr.ActivityName, &sr.JumpURL,
			&sr.RecommendType, &sr.IsShow, &sr.Sort,
			&sr.ShareTitle, &sr.ShareIconURL, &sr.ShareURL, &sr.ShareDesc)

		srs = append(srs, sr)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// AddSearchRecommend 新增搜索推荐
func (c *Copywriter) AddSearchRecommend(sr SearchRecommend) (id int64, err error) {
	insertSQL := utils.SetSQLFormat(`
		INSERT INTO gpxj_app.t_search_recommend (
			stock_name,stock_code,activity_name,jump_url,
			share_title,share_desc,share_icon_url,share_url,
			recommend_type,is_show,sort) 
		VALUES
			(
			'{0}','{1}','{2}','{3}',
			'{4}','{5}','{6}','{7}',
			'{8}','{9}','{10}');
		`, sr.StockName, sr.StockCode, sr.ActivityName, sr.JumpURL,
		sr.ShareTitle, sr.ShareDesc, sr.ShareIconURL, sr.ShareURL,
		sr.RecommendType, sr.IsShow, sr.Sort)

	glog.Info("insertSQL:", insertSQL)

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

// ModifySearchRecommend 修改搜索推荐
func (c *Copywriter) ModifySearchRecommend(sr SearchRecommend) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
	UPDATE 
		gpxj_app.t_search_recommend 
	SET
		stock_name = '{1}',
		stock_code = '{2}',
		activity_name = '{3}',
		jump_url = '{4}',
		share_title = '{5}',
		share_desc = '{6}',
		share_icon_url = '{7}',
		share_url = '{8}',
		recommend_type = '{9}',
		is_show = '{10}',
		sort = '{11}'
	WHERE id = '{0}' ;
	`, sr.ID,
		sr.StockName, sr.StockCode, sr.ActivityName, sr.JumpURL,
		sr.ShareTitle, sr.ShareDesc, sr.ShareIconURL, sr.ShareURL,
		sr.RecommendType, sr.IsShow, sr.Sort)

	glog.Info("updateSQL:", updateSQL)

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

// DropSearchRecommend 删除搜索推荐
func (c *Copywriter) DropSearchRecommend(sr SearchRecommend) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_search_recommend WHERE id ='{0}'`, sr.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}
	return
}

// MarketingLabel 行情营销标签
type MarketingLabel struct {
	ID           int    `json:"id"`
	TagName      string `json:"tag_name"`
	JumpURL      string `json:"jump_url"`
	Place        string `json:"place"`
	IsShow       int    `json:"is_show"`
	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
}

// GetMarketingLabels  查询行情营销标签
func (c *Copywriter) GetMarketingLabels(status string) (labels []MarketingLabel, err error) {
	labels = make([]MarketingLabel, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,tag_name,jump_url,place,is_show,
				share_title,share_desc,share_icon_url,share_url
		FROM gpxj_app.t_marketing_label 
 		WHERE 1
		`)

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 0`, querySQL)
		}
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var label MarketingLabel
		rows.Scan(&label.ID, &label.TagName, &label.JumpURL, &label.Place, &label.IsShow,
			&label.ShareTitle, &label.ShareDesc, &label.ShareIconURL, &label.ShareURL)

		labels = append(labels, label)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// AddMarketingLabel 新增行情营销标签
func (c *Copywriter) AddMarketingLabel(ml MarketingLabel) (id int64, err error) {
	insertSQL := utils.SetSQLFormat(`
	INSERT INTO gpxj_app.t_marketing_label (
		tag_name,jump_url,place,is_show,
		share_title,share_desc,share_icon_url,share_url) 
	  VALUES
		(
		  '{0}','{1}','{2}','{3}',
		  '{4}','{5}','{6}','{7}'
		);
		`, ml.TagName, ml.JumpURL, ml.Place, ml.IsShow,
		ml.ShareTitle, ml.ShareDesc, ml.ShareIconURL, ml.ShareURL)
	glog.Info("insertSQL:", insertSQL)

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

// ModifyMarketingLabel 修改行情营销标签
func (c *Copywriter) ModifyMarketingLabel(ml MarketingLabel) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
	UPDATE 
		gpxj_app.t_marketing_label 
	SET
		tag_name = '{1}',
		jump_url = '{2}',
		place = '{3}',
		is_show = '{4}',
		share_title = '{5}',
		share_desc = '{6}',
		share_icon_url = '{7}',
		share_url = '{8}'
	WHERE id = '{0}';
	`, ml.ID,
		ml.TagName, ml.JumpURL, ml.Place, ml.IsShow,
		ml.ShareTitle, ml.ShareDesc, ml.ShareIconURL, ml.ShareURL)
	glog.Info("updateSQL:", updateSQL)

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

// DropMarketingLabel 删除行情营销标签
func (c *Copywriter) DropMarketingLabel(ml MarketingLabel) (id int64, err error) {
	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_marketing_label WHERE id ='{0}'`, ml.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}
	return
}
