package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫莱夫里耶MaulevrierBarony struct {
	feud.BaseBarony
}

var BMaulevrier莫莱夫里耶 feud.Barony = &莫莱夫里耶MaulevrierBarony{}

func init() {
    f := BMaulevrier莫莱夫里耶.(*莫莱夫里耶MaulevrierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maulevrier",
		TitleName: "莫莱夫里耶",
		TitleCode: "b_maulevrier",
	}
}
