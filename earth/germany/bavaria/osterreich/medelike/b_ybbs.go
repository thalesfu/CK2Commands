package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊布斯YbbsBarony struct {
	feud.BaseBarony
}

var BYbbs伊布斯 feud.Barony = &伊布斯YbbsBarony{}

func init() {
    f := BYbbs伊布斯.(*伊布斯YbbsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ybbs",
		TitleName: "伊布斯",
		TitleCode: "b_ybbs",
	}
}
