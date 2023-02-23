package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞雷斯SeresBarony struct {
	feud.BaseBarony
}

var BSeres塞雷斯 feud.Barony = &塞雷斯SeresBarony{}

func init() {
	f := BSeres塞雷斯.(*塞雷斯SeresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seres",
		TitleName: "塞雷斯",
		TitleCode: "b_seres",
	}
}
