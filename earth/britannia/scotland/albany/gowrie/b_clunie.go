package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克卢尼ClunieBarony struct {
	feud.BaseBarony
}

var BClunie克卢尼 feud.Barony = &克卢尼ClunieBarony{}

func init() {
	f := BClunie克卢尼.(*克卢尼ClunieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clunie",
		TitleName: "克卢尼",
		TitleCode: "b_clunie",
	}
}
