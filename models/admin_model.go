package models

import (
	"iris_test/common"
)

type Admin struct {
	Id  int       `gorm:"primary_key;unsigned;comment:'主键id'" json:"id"`
	Username string `gorm:"type:VARCHAR(20);not null;default:'';comment:'用户名'" json:"username"`
	Password string `gorm:"type:VARCHAR(50);not null;default:'';comment:'密码'" json:"password"`
	Status    uint8   `gorm:"type:tinyint;unsigned;not null;default:1;comment:'状态，1正常，2禁用'" json:"status"`
}

//详情
func GetAdminByUserName(username string, password string) Admin {
	password = common.Md5(password)
	info := Admin{}
	db.Model(&Admin{}).Where("username = ?", username).Where("password = ?", password).First(&info)
	return info
}

//详情
func GetAdmin(id int) Admin {
	info := Admin{}
	db.Model(&Admin{}).Where("id = ?", id).First(&info)
	return info
}

//列表
func GetAdminList(page int, per_page int) []Admin {
	offset := (page-1)*per_page
	list := []Admin{}
	db.Model(&Admin{}).Limit(per_page).Offset(offset).Order("id asc", true).Find(&list)
	return list
}

//列表
func GetAdminListPaginator(page int) ([]Admin, int64) {
	per_page := PAGE_LIMIT
	offset := (page-1)*per_page
	list := []Admin{}
	db.Model(&Admin{}).Limit(per_page).Offset(offset).Order("id asc", true).Find(&list)

	var nums int64 = 0
	db.Model(&Admin{}).Count(&nums)

	return list, nums
}

//新增
func CreateAdmin(info Admin) Admin {
	info.Password = common.Md5(info.Password)
	db.Model(&Admin{}).Create(&info)
	return info
}

//更新
func UpdateAdmin(info Admin)Admin {
	info.Password = common.Md5(info.Password)
	db.Model(&Admin{}).Save(&info)
	return info
}

//更新
func UpdateAdminInfo(info Admin)Admin {
	db.Model(&Admin{}).Save(&info)
	return info
}

//删除
func DeleteAdmin(id int)Admin {
	info := Admin{}
	info.Id = id
	db.Model(&Admin{}).Delete(&info)
	return info
}