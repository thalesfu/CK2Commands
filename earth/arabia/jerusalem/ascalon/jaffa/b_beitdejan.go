package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝特达甘BeitdejanBarony struct {
	feud.BaseBarony
}

var BBeitdejan贝特达甘 feud.Barony = &贝特达甘BeitdejanBarony{}

func init() {
	f := BBeitdejan贝特达甘.(*贝特达甘BeitdejanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beitdejan",
		TitleName: "贝特达甘",
		TitleCode: "b_beitdejan",
	}
}
