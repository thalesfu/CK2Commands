package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代胡克DayhoukBarony struct {
	feud.BaseBarony
}

var BDayhouk代胡克 feud.Barony = &代胡克DayhoukBarony{}

func init() {
	f := BDayhouk代胡克.(*代胡克DayhoukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dayhouk",
		TitleName: "代胡克",
		TitleCode: "b_dayhouk",
	}
}
