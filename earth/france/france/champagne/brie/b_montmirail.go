package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙米拉伊MontmirailBarony struct {
	feud.BaseBarony
}

var BMontmirail蒙米拉伊 feud.Barony = &蒙米拉伊MontmirailBarony{}

func init() {
	f := BMontmirail蒙米拉伊.(*蒙米拉伊MontmirailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montmirail",
		TitleName: "蒙米拉伊",
		TitleCode: "b_montmirail",
	}
}
