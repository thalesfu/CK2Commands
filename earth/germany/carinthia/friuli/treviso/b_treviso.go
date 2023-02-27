package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷维索TrevisoBarony struct {
	feud.BaseBarony
}

var BTreviso特雷维索 feud.Barony = &特雷维索TrevisoBarony{}

func init() {
    f := BTreviso特雷维索.(*特雷维索TrevisoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treviso",
		TitleName: "特雷维索",
		TitleCode: "b_treviso",
	}
}
