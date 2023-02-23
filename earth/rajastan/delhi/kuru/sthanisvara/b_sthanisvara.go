package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨他泥湿伐罗SthanisvaraBarony struct {
	feud.BaseBarony
}

var BSthanisvara萨他泥湿伐罗 feud.Barony = &萨他泥湿伐罗SthanisvaraBarony{}

func init() {
	f := BSthanisvara萨他泥湿伐罗.(*萨他泥湿伐罗SthanisvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sthanisvara",
		TitleName: "萨他泥湿伐罗",
		TitleCode: "b_sthanisvara",
	}
}
