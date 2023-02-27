package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切沙ChoshaBarony struct {
	feud.BaseBarony
}

var BChosha切沙 feud.Barony = &切沙ChoshaBarony{}

func init() {
    f := BChosha切沙.(*切沙ChoshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chosha",
		TitleName: "切沙",
		TitleCode: "b_chosha",
	}
}
