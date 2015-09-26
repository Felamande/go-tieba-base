package apis

import (
	"encoding/json"

	"github.com/purstal/go-tieba-base/simple-http"
	"github.com/purstal/go-tieba-base/tieba"
)

func CancelBlockIDWeb(acc *postbar.Account,
	forumName, userID,
	userName string) (error, *postbar.PbError) {

	var parameters simple_http.Parameters
	parameters.Add("word", forumName)

	tbs, err, pberr := GetTbs(acc)
	if err != nil {
		return err, nil
	} else if pberr != nil {
		return nil, pberr
	}
	parameters.Add("tbs", tbs)

	parameters.Add("ie", "gbk")
	parameters.Add("type", "0")
	parameters.Add("list[0][user_id]", userID)
	parameters.Add("list[0][user_name]", userName)

	var cookies simple_http.Cookies
	cookies.Add("BDUSS", acc.BDUSS)

	resp, err := simple_http.Get("http://tieba.baidu.com/bawu2/platform/cancelFilter", parameters, cookies)

	if err != nil {
		return err, nil
	}
	var x struct {
		ErrorCode int    `json:"errno"`
		ErrorMsg  string `json:"errmsg"`
	}

	json.Unmarshal(resp, &x)

	return nil, postbar.NewPbError(x.ErrorCode, x.ErrorMsg)

}
