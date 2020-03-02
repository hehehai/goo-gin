package models

type Article struct {
	Model

	TagId int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId:      data["tag_id"].(int),
		Title:      data["title"].(string),
		Desc:       data["desc"].(string),
		Content:    data["content"].(string),
		CreatedBy:  data["created_by"].(string),
		State:      data["state"].(int),
	})

	return true
}

//判断文章id是否存在
func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).Find(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// 文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

//获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//获取单个文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).Find(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

//编辑文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})

	return true
}