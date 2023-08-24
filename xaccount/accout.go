/*
功能：帐号
说明：
*/
package xaccount

import "regexp"

// Check_account 校验帐号
func Check_account(account string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]{6,16}$", account)
	return match
}

// Check_nickname 校验昵称
func Check_nickname(nickname string) bool {
	return true
}

// Check_pwd 校验密码
func Check_pwd(pwd string) bool {
	match, _ := regexp.MatchString("^([0-9]|[a-zA-z]|[0-9A-Za-z]){6,16}$", pwd)
	//match, _ := regexp.MatchString("(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,50}$", pwd)
	return match
}

// Check_phone 校验手机号
func Check_phone(phone string) bool {
	regular := `^1[3-9]\d{9}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
