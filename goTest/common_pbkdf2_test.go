package goTest

import (
	"fmt"
	"github.com/daida459031925/common/encryption/pbkdf2"
	err "github.com/daida459031925/common/error"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestPbkdf2(t *testing.T) {
	pwd, salt, e := pbkdf2.EncryptPwd("pwd string")
	logx.Info(pwd)
	logx.Info(salt)
	logx.Info(e)
	if e != nil {
		err.New("加密失败")
	}
	fmt.Println(pbkdf2.CheckEncryptPwdMatch(pwd, "pwd string"))
	fmt.Println(pbkdf2.CheckEncryptPwdMatch(pwd, "pwdstring"))
	fmt.Println(pbkdf2.CheckEncryptPwdMatch("pwd string", pwd))

}
