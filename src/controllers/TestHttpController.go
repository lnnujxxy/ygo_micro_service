package controllers

import (
	"fmt"

	"github.com/mlaoji/yclient"
	"github.com/mlaoji/ygo/controllers"
	"github.com/mlaoji/ygo/lib"
)

type TestHttpController struct {
	controllers.BaseController
}

func (this *TestHttpController) HelloAction() { // {{{

	msg := this.GetString("msg")

	lib.Interceptor(len(msg) > 0, lib.ERR_PARAMS, "msg")

	ret := map[string]interface{}{
		"msg": msg,
	}

	this.Render(ret)

}

func (this *TestHttpController) RpcAction() {
	c, err := yclient.NewYClient("127.0.0.1:6001", "passport", "ep8oQ8a87AfN")
	if err != nil {
		panic("rpc connect fail " + err.Error())
	}

	data, error := c.Request("testRpc/Hello", map[string]string{"msg": "123123"})

	if error != nil {
		fmt.Printf("%#v", error)

		//获取错误码
		errno := c.Errno(error)
		fmt.Printf("%#v", errno)

	} else {

		fmt.Println("ok")
		fmt.Printf("%#v", data)
	}
}
