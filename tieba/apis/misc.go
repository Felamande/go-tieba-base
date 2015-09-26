package apis

import (
	"encoding/json"

	"github.com/purstal/go-tieba-base/simple-http"
	"github.com/purstal/go-tieba-base/tieba"
)

func RSearchForum(forumName string) ([]byte, error) {
	var parameters simple_http.Parameters
	parameters.Add("query", forumName)
	postbar.AddSignature(&parameters)
	return simple_http.Post(`http://c.tieba.baidu.com/c/f/forum/search`, parameters)
}

type ForumSearchResult struct {
	ForumID   uint64 `json:"forum_id"`
	ForumName string `json:"forum_name"`
}

func SearchForum(forumName string) ([]ForumSearchResult, error, *postbar.PbError) {
	resp, err := RSearchForum(forumName)
	if err != nil {
		return nil, err, nil
	}

	var forumSearchResults struct {
		ForumList []ForumSearchResult `json:"forum_list"`
		ErrorCode int                 `json:"error_code,string"`
		ErrorMsg  string              `json:"error_msg"`
	}

	json.Unmarshal(resp, &forumSearchResults)

	if forumSearchResults.ErrorCode == 110003 {
		return nil, nil, nil
	} else if forumSearchResults.ErrorCode != 0 {
		return nil, nil, postbar.NewPbError(forumSearchResults.ErrorCode, forumSearchResults.ErrorMsg)
	}

	return forumSearchResults.ForumList, nil, nil

}
