package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加摩GarmoBarony struct {
	feud.BaseBarony
}

var BGarmo加摩 feud.Barony = &加摩GarmoBarony{}

func init() {
    f := BGarmo加摩.(*加摩GarmoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garmo",
		TitleName: "加摩",
		TitleCode: "b_garmo",
	}
}
