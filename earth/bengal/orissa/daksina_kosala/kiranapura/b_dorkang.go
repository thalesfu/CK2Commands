package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尔康DorkangBarony struct {
	feud.BaseBarony
}

var BDorkang多尔康 feud.Barony = &多尔康DorkangBarony{}

func init() {
    f := BDorkang多尔康.(*多尔康DorkangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorkang",
		TitleName: "多尔康",
		TitleCode: "b_dorkang",
	}
}
