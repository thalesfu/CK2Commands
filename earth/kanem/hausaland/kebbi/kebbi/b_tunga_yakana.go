package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通加亚卡纳Tunga_yakanaBarony struct {
	feud.BaseBarony
}

var BTunga_yakana通加亚卡纳 feud.Barony = &通加亚卡纳Tunga_yakanaBarony{}

func init() {
    f := BTunga_yakana通加亚卡纳.(*通加亚卡纳Tunga_yakanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tunga_yakana",
		TitleName: "通加亚卡纳",
		TitleCode: "b_tunga_yakana",
	}
}
