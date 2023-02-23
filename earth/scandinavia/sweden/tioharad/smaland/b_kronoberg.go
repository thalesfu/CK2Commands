package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁努贝里KronobergBarony struct {
	feud.BaseBarony
}

var BKronoberg克鲁努贝里 feud.Barony = &克鲁努贝里KronobergBarony{}

func init() {
	f := BKronoberg克鲁努贝里.(*克鲁努贝里KronobergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kronoberg",
		TitleName: "克鲁努贝里",
		TitleCode: "b_kronoberg",
	}
}
