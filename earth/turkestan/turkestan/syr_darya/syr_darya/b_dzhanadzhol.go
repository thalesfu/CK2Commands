package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾纳_焦尔DzhanadzholBarony struct {
	feud.BaseBarony
}

var BDzhanadzhol贾纳_焦尔 feud.Barony = &贾纳_焦尔DzhanadzholBarony{}

func init() {
    f := BDzhanadzhol贾纳_焦尔.(*贾纳_焦尔DzhanadzholBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhanadzhol",
		TitleName: "贾纳_焦尔",
		TitleCode: "b_dzhanadzhol",
	}
}
