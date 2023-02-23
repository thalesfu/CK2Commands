package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苇关WeiguanBarony struct {
	feud.BaseBarony
}

var BWeiguan苇关 feud.Barony = &苇关WeiguanBarony{}

func init() {
	f := BWeiguan苇关.(*苇关WeiguanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weiguan",
		TitleName: "苇关",
		TitleCode: "b_weiguan",
	}
}
