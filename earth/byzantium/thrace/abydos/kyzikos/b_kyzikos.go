package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基齐库斯KyzikosBarony struct {
	feud.BaseBarony
}

var BKyzikos基齐库斯 feud.Barony = &基齐库斯KyzikosBarony{}

func init() {
	f := BKyzikos基齐库斯.(*基齐库斯KyzikosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzikos",
		TitleName: "基齐库斯",
		TitleCode: "b_kyzikos",
	}
}
