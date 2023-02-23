package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴德比BodbeBarony struct {
	feud.BaseBarony
}

var BBodbe巴德比 feud.Barony = &巴德比BodbeBarony{}

func init() {
	f := BBodbe巴德比.(*巴德比BodbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodbe",
		TitleName: "巴德比",
		TitleCode: "b_bodbe",
	}
}
