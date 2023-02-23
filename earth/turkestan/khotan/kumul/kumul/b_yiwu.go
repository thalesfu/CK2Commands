package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊吾YiwuBarony struct {
	feud.BaseBarony
}

var BYiwu伊吾 feud.Barony = &伊吾YiwuBarony{}

func init() {
	f := BYiwu伊吾.(*伊吾YiwuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yiwu",
		TitleName: "伊吾",
		TitleCode: "b_yiwu",
	}
}
