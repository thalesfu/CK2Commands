package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科帕尔贝里KopparbergBarony struct {
	feud.BaseBarony
}

var BKopparberg科帕尔贝里 feud.Barony = &科帕尔贝里KopparbergBarony{}

func init() {
	f := BKopparberg科帕尔贝里.(*科帕尔贝里KopparbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kopparberg",
		TitleName: "科帕尔贝里",
		TitleCode: "b_kopparberg",
	}
}
