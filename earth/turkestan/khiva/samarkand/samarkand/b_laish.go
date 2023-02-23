package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱什LaishBarony struct {
	feud.BaseBarony
}

var BLaish莱什 feud.Barony = &莱什LaishBarony{}

func init() {
	f := BLaish莱什.(*莱什LaishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laish",
		TitleName: "莱什",
		TitleCode: "b_laish",
	}
}
