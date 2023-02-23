package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿灵索斯AringsasBarony struct {
	feud.BaseBarony
}

var BAringsas阿灵索斯 feud.Barony = &阿灵索斯AringsasBarony{}

func init() {
	f := BAringsas阿灵索斯.(*阿灵索斯AringsasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aringsas",
		TitleName: "阿灵索斯",
		TitleCode: "b_aringsas",
	}
}
