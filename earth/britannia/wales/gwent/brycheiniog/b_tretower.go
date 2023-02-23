package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷陶尔TretowerBarony struct {
	feud.BaseBarony
}

var BTretower特雷陶尔 feud.Barony = &特雷陶尔TretowerBarony{}

func init() {
	f := BTretower特雷陶尔.(*特雷陶尔TretowerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tretower",
		TitleName: "特雷陶尔",
		TitleCode: "b_tretower",
	}
}
