package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴萨罗洛BasaraluBarony struct {
	feud.BaseBarony
}

var BBasaralu巴萨罗洛 feud.Barony = &巴萨罗洛BasaraluBarony{}

func init() {
	f := BBasaralu巴萨罗洛.(*巴萨罗洛BasaraluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basaralu",
		TitleName: "巴萨罗洛",
		TitleCode: "b_basaralu",
	}
}
