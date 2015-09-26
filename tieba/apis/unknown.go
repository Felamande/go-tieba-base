package apis

import (
	"github.com/purstal/go-tieba-base/misc"
	"github.com/purstal/go-tieba-base/simple-http"
	"github.com/purstal/go-tieba-base/tieba"
)

func RTest(acc *postbar.Account, userName, password string) (
	[]byte, error) {

	var parameters simple_http.Parameters
	parameters.Add("un", userName)
	parameters.Add("passwd", misc.ComputeBase64(password))
	postbar.ProcessParams(&parameters, acc)
	return simple_http.Post("http://c.tieba.baidu.com/c/s/test", parameters)
}
