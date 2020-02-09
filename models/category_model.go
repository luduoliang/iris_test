package models

type Category struct {
	Id  int       `gorm:"primary_key;unsigned;comment:'主键id'" json:"id"`
	Name string `gorm:"type:VARCHAR(50);not null;default:'';comment:'分类名称'" json:"name"`
	Weight int `gorm:"type:int;not null;default:0;comment:'权重（升序）'" json:"weight"`
}


//详情
func GetCategory(id int) Category {
	info := Category{}
	db.Model(&Category{}).Where("id = ?", id).First(&info)
	return info
}

//列表
func GetCategoryList(page int, per_page int) []Category {
	offset := (page-1)*per_page
	list := []Category{}
	db.Model(&Category{}).Limit(per_page).Offset(offset).Order("id asc", true).Find(&list)
	return list
}

//列表
func GetAllCategoryList() []Category {
	list := []Category{}
	db.Model(&Category{}).Order("id asc", true).Find(&list)
	return list
}



//列表
func GetCategoryListPaginator(page int) ([]Category, int64) {
	per_page := PAGE_LIMIT
	offset := (page-1)*per_page
	list := []Category{}
	db.Model(&Category{}).Limit(per_page).Offset(offset).Order("id asc", true).Find(&list)

	var nums int64 = 0
	db.Model(&Category{}).Count(&nums)

	return list, nums
}

//新增
func CreateCategory(info Category) Category {
	db.Model(&Category{}).Create(&info)
	return info
}

//更新
func UpdateCategory(info Category)Category {
	db.Model(&Category{}).Save(&info)
	return info
}


//删除
func DeleteCategory(id int)Category {
	info := Category{}
	info.Id = id
	db.Model(&Category{}).Delete(&info)
	return info
}