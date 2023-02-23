package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内奥利亚尼NeoljaniBarony struct {
	feud.BaseBarony
}

var BNeoljani内奥利亚尼 feud.Barony = &内奥利亚尼NeoljaniBarony{}

func init() {
	f := BNeoljani内奥利亚尼.(*内奥利亚尼NeoljaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neoljani",
		TitleName: "内奥利亚尼",
		TitleCode: "b_neoljani",
	}
}
