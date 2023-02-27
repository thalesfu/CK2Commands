package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾哈迈达巴德AhmadabadBarony struct {
	feud.BaseBarony
}

var BAhmadabad艾哈迈达巴德 feud.Barony = &艾哈迈达巴德AhmadabadBarony{}

func init() {
    f := BAhmadabad艾哈迈达巴德.(*艾哈迈达巴德AhmadabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahmadabad",
		TitleName: "艾哈迈达巴德",
		TitleCode: "b_ahmadabad",
	}
}
