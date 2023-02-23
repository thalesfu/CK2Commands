package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫泽拉达尔HezardarBarony struct {
	feud.BaseBarony
}

var BHezardar赫泽拉达尔 feud.Barony = &赫泽拉达尔HezardarBarony{}

func init() {
	f := BHezardar赫泽拉达尔.(*赫泽拉达尔HezardarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hezardar",
		TitleName: "赫泽拉达尔",
		TitleCode: "b_hezardar",
	}
}
