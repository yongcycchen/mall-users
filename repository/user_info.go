package repository

import (
	"gitee.com/kelvins-io/kelvins"
	"github.com/yongcycchen/mall-users/model/mysql"
)


func CreateUser(user *mysql.User) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Insert(user)
	return
}

func FindUserInfo(sqlSelect string, uidList []int64) ([]mysql.User, error) {
	var result = make([]mysql.User, 0)
	err := kelvins.XORM_DBEngine.Table(mysql.TableUser).Select(sqlSelect).In("id", uidList).Find(&result)
	return result, err
}

func GetUserByUid(uid int) (*mysql.User, error) {
	var user mysql.User
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Where("id = ?", uid).Get(&user)
	return &user, err
}

func UpdateUserInfo(query, maps interface{}) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Where(query).Update(maps)
	return
}

func GetUserByEmail(sqlSelect, email string) (*mysql.User, error) {
	var user mysql.User
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Select(sqlSelect).Where("email = ? ", email).Get(&user)
	return &user, err
}

func GetUserByPhone(sqlSelect, countryCode, phone string) (*mysql.User, error) {
	var user mysql.User
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Select(sqlSelect).Where("country_code = ? and phone = ?", countryCode, phone).Get(&user)
	return &user, err
}

func GetUserByInviteCode(inviteCode string) (*mysql.User, error) {
	var user mysql.User
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Where("invite_code = ?", inviteCode).Get(&user)
	return &user, err
}

func CheckUserExistById(id int) (exist bool, err error) {
	var user mysql.User
	exist, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).
		Select("id").
		Where("id = ?", id).Get(&user)
	if err != nil {
		return false, err
	}
	if user.Id != 0 {
		return true, nil
	}

	return false, nil
}

func CheckUserExistByPhone(countryCode, phone string) (exist bool, err error) {
	var user mysql.User
	exist, err = kelvins.XORM_DBEngine.Table(mysql.TableUser).
		Select("id").
		Where("country_code = ? and phone = ?", countryCode, phone).Get(&user)
	if err != nil {
		return false, err
	}
	if user.Id > 0 {
		return true, nil
	}

	return false, nil
}

func ListUserInfo(sqlSelect string, pageSize, pageNum int) (result []mysql.User, err error) {
	result = make([]mysql.User, 0)
	err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Select(sqlSelect).
		Limit(pageSize, (pageNum-1)*pageSize).
		Find(&result)
	return result, err
}
