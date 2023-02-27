package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔博尔MnttaborBarony struct {
	feud.BaseBarony
}

var BMnttabor塔博尔 feud.Barony = &塔博尔MnttaborBarony{}

func init() {
    f := BMnttabor塔博尔.(*塔博尔MnttaborBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mnttabor",
		TitleName: "塔博尔",
		TitleCode: "b_mnttabor",
	}
}
