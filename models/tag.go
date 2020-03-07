package models

//tag 表结构
type Tag struct {
	//基础结构
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//标签列表 页，页大小，查询参数 => 标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	//查询 => 偏移 => 分页 => 赋值
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

//标签总数 查询参数 => int
func GetTagTotal(maps interface{}) (count int) {
	//选择表 => 查询 => 总数
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

//判断标签名是否已存在 标签名 => bool
func ExistTagByName(name string) bool {
	var tag Tag
	//查询 => 赋值
	db.Select("id").Where("name = ?", name).First(&tag)
	//是否存在
	if tag.ID > 0 {
		return true
	}

	return false
}

//判断标签是否存在 标签id => bool
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).Find(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//添加标签 标签名, 状态，创建者
func AddTag(name string, state int, createdBy string) bool {
	// 创建
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

//更新标签数据 id, data => bool
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}

//删除标签 id => bool
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// 物理删除以及软删除的 tag
func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}