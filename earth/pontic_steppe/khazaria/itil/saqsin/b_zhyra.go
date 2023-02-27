package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日拉ZhyraBarony struct {
	feud.BaseBarony
}

var BZhyra日拉 feud.Barony = &日拉ZhyraBarony{}

func init() {
    f := BZhyra日拉.(*日拉ZhyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhyra",
		TitleName: "日拉",
		TitleCode: "b_zhyra",
	}
}
