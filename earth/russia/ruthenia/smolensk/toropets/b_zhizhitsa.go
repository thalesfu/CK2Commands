package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日日察ZhizhitsaBarony struct {
	feud.BaseBarony
}

var BZhizhitsa日日察 feud.Barony = &日日察ZhizhitsaBarony{}

func init() {
    f := BZhizhitsa日日察.(*日日察ZhizhitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhizhitsa",
		TitleName: "日日察",
		TitleCode: "b_zhizhitsa",
	}
}
