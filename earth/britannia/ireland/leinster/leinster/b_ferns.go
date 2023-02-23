package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗恩斯FernsBarony struct {
	feud.BaseBarony
}

var BFerns弗恩斯 feud.Barony = &弗恩斯FernsBarony{}

func init() {
	f := BFerns弗恩斯.(*弗恩斯FernsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferns",
		TitleName: "弗恩斯",
		TitleCode: "b_ferns",
	}
}
