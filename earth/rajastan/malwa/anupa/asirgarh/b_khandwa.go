package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯德瓦KhandwaBarony struct {
	feud.BaseBarony
}

var BKhandwa肯德瓦 feud.Barony = &肯德瓦KhandwaBarony{}

func init() {
    f := BKhandwa肯德瓦.(*肯德瓦KhandwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khandwa",
		TitleName: "肯德瓦",
		TitleCode: "b_khandwa",
	}
}
