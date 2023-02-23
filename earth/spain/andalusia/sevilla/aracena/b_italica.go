package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊塔利卡ItalicaBarony struct {
	feud.BaseBarony
}

var BItalica伊塔利卡 feud.Barony = &伊塔利卡ItalicaBarony{}

func init() {
	f := BItalica伊塔利卡.(*伊塔利卡ItalicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "italica",
		TitleName: "伊塔利卡",
		TitleCode: "b_italica",
	}
}
