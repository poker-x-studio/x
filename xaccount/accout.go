/*
功能：帐号
说明：
特殊需求表达式 https://c.runoob.com/front-end/854/
*/
package xaccount

import (
	"fmt"
	"regexp"
)

// Check_account 校验帐号
func Check_account(account string, min_len int, max_len int) (bool, error) {
	len := len(account)
	if len < min_len {
		return false, fmt.Errorf("错误,账号最小长度为%d", min_len)
	}
	if len > max_len {
		return false, fmt.Errorf("错误,账号最大长度为%d", max_len)
	}
	pattern := fmt.Sprintf("^[0-9A-Za-z_]{%d,%d}$", min_len, max_len)
	matched, _ := regexp.MatchString(pattern, account)
	if !matched {
		return false, fmt.Errorf("错误,账号不合乎当前规则")
	}
	return matched, nil
}

// Check_nickname 校验昵称
func Check_nickname(nickname string) bool {
	return true
}

// Check_pwd 校验密码
func Check_pwd(pwd string, min_len int, max_len int) (bool, error) {
	len := len(pwd)
	if len < min_len {
		return false, fmt.Errorf("错误,密码最小长度为%d", min_len)
	}
	if len > max_len {
		return false, fmt.Errorf("错误,密码最大长度为%d", max_len)
	}
	pattern := fmt.Sprintf("^[0-9A-Za-z_]{%d,%d}$", min_len, max_len)
	matched, _ := regexp.MatchString(pattern, pwd)
	if !matched {
		return false, fmt.Errorf("错误,密码不合乎当前规则")
	}
	return matched, nil
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
