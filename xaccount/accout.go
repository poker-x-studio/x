/*
功能：帐号
说明：
特殊需求表达式 https://c.runoob.com/front-end/854/
*/
package xaccount

import "regexp"

// Check_account 校验帐号
func Check_account(account string) bool {
	is_matched, _ := regexp.MatchString("^[0-9A-Za-z_]{6,16}$", account)
	return is_matched
}

// Check_nickname 校验昵称
func Check_nickname(nickname string) bool {
	return true
}

// Check_pwd 校验密码
func Check_pwd(pwd string) bool {
	is_matched, _ := regexp.MatchString("^[0-9A-Za-z_]{6,16}$", pwd)
	return is_matched
}

// Check_phone 校验手机号
func Check_phone(phone string) bool {
	is_matched, _ := regexp.MatchString("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", phone)
	return is_matched
}

// Check_email 校验email
func Check_email(email string) bool {
	is_matched, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", email)
	return is_matched
}

// Check_id 校验身份证
func Check_id(id string) bool {
	is_matched, _ := regexp.MatchString("(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)", id)
	return is_matched
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
