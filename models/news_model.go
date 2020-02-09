package models

type News struct {
	Id  int       `gorm:"primary_key;unsigned;comment:'主键id'" json:"id"`
	CategoryId int `gorm:"type:int;not null;default:0;comment:'分类id'" json:"category_id"`
	Title string `gorm:"type:VARCHAR(250);not null;default:'';comment:'标题'" json:"title"`
	Description string `gorm:"type:VARCHAR(500);not null;default:'';comment:'简介'" json:"description"`
	Pic string `gorm:"type:VARCHAR(200);not null;default:'';comment:'图片'" json:"pic"`
	Content string `gorm:"type:text;not null;default:'';comment:'内容'" json:"title"`
	Weight int `gorm:"type:int;not null;default:0;comment:'权重（升序）'" json:"weight"`

	Category Category `gorm:"foreignkey:CategoryId" json:"category"` //关联Category表，用Preload("Category")查询时会带上分类信息
}


//详情
func GetNews(id int) News {
	info := News{}
	db.Model(&News{}).Where("id = ?", id).First(&info)
	return info
}

//详情
func GetNewsPreload(id int) News {
	info := News{}
	db.Model(&News{}).Preload("Category").Where("id = ?", id).First(&info)
	return info
}

//列表
func GetNewsList(category_id int, page int, per_page int) []News {
	offset := (page-1)*per_page
	list := []News{}
	query := db.Model(&News{}).Preload("Category")
	if(category_id > 0){
		query = query.Where("category_id = ?", category_id)
	}
	query.Limit(per_page).Offset(offset).Order("id desc", true).Find(&list)
	return list
}

//列表
func GetNewsListPaginator(page int) ([]News, int64) {
	per_page := PAGE_LIMIT
	offset := (page-1)*per_page
	list := []News{}
	db.Model(&News{}).Preload("Category").Limit(per_page).Offset(offset).Order("id desc", true).Find(&list)

	var nums int64 = 0
	db.Model(&News{}).Count(&nums)

	return list, nums
}

//新增
func CreateNews(info News) News {
	db.Model(&News{}).Create(&info)
	return info
}

//更新
func UpdateNews(info News)News {
	db.Model(&News{}).Save(&info)
	return info
}

//删除
func DeleteNews(id int)News {
	info := News{}
	info.Id = id
	db.Model(&News{}).Delete(&info)
	return info
}