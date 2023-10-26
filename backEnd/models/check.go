package models
import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
)

func Check(ctx *context.Context){
	fmt.Println(ctx.Input.Method())
	fmt.Println(ctx.Input.URL())
	if ctx.Input.Method() == "OPTIONS" {
		return
	}

	auth := ctx.Input.Header("Authorization")
	if strings.Compare(auth, "") == 0 {
		jsonbyte, _ := json.Marshal(GetAnswerByErrorCode(RESULT_TOKEN_ERROR))
		ctx.ResponseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
		ctx.ResponseWriter.Write(jsonbyte)
		ctx.ResponseWriter.Flush()
		return
	}

	if err != RESULT_OK {
		jsonbyte, _ := json.Marshal(GetAnswerByErrorCode(err))
		ctx.ResponseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
		ctx.ResponseWriter.Write(jsonbyte)
		ctx.ResponseWriter.Flush()
		return
	}

	ctx.Input.SetData("userinfo", user)
}



