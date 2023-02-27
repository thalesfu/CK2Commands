package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱内LaaneBarony struct {
	feud.BaseBarony
}

var BLaane莱内 feud.Barony = &莱内LaaneBarony{}

func init() {
    f := BLaane莱内.(*莱内LaaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laane",
		TitleName: "莱内",
		TitleCode: "b_laane",
	}
}
