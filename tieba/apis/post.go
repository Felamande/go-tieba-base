package apis

import (
	"strconv"
	"time"

	"github.com/purstal/go-tieba-base/simple-http"
	"github.com/purstal/go-tieba-base/tieba"
)

func _AddPost(acc *postbar.Account, content string, fid uint64, forumName string, tid uint64) {

}

func AddPost(accAndr *postbar.Account, content string, fid uint64, forumName string, tid uint64, floorNumber int, quoteID uint64) (error, *postbar.PbError) {
	var parameters simple_http.Parameters
	parameters.Add("anonymous", "1")
	parameters.Add("content", content)
	parameters.Add("cuid", postbar.GenCUID(accAndr.ClientID, accAndr.PhoneIMEI))
	parameters.Add("fid", strconv.FormatUint(fid, 10))
	parameters.Add("floor_num", "0")
	parameters.Add("is_ad", "0")
	parameters.Add("is_addition", "0")
	parameters.Add("kw", forumName)
	parameters.Add("new_vcode", "1")
	if quoteID != 0 {
		parameters.Add("quote_id", strconv.FormatUint(quoteID, 10))
	}
	for {
		tbs, err := GetTbsWeb(accAndr.BDUSS)
		if err == nil {
			parameters.Add("tbs", tbs)
			break
		}
	}
	parameters.Add("tid", strconv.FormatUint(tid, 10))
	parameters.Add("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	parameters.Add("vcode_tag", "11")

	postbar.ProcessParams(&parameters, accAndr)

	println(parameters.Encode())

	resp, err := simple_http.Post("http://c.tieba.baidu.com/c/c/post/add", parameters)

	println(string(resp))

	if err != nil {
		return err, nil
	}

	return nil, nil

}
