package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列城LehBarony struct {
	feud.BaseBarony
}

var BLeh列城 feud.Barony = &列城LehBarony{}

func init() {
	f := BLeh列城.(*列城LehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leh",
		TitleName: "列城",
		TitleCode: "b_leh",
	}
}
