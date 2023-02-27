package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米利波波利斯MilitopolisBarony struct {
	feud.BaseBarony
}

var BMilitopolis米利波波利斯 feud.Barony = &米利波波利斯MilitopolisBarony{}

func init() {
    f := BMilitopolis米利波波利斯.(*米利波波利斯MilitopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "militopolis",
		TitleName: "米利波波利斯",
		TitleCode: "b_militopolis",
	}
}
