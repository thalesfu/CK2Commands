package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚祖尔YazurBarony struct {
	feud.BaseBarony
}

var BYazur亚祖尔 feud.Barony = &亚祖尔YazurBarony{}

func init() {
	f := BYazur亚祖尔.(*亚祖尔YazurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yazur",
		TitleName: "亚祖尔",
		TitleCode: "b_yazur",
	}
}
