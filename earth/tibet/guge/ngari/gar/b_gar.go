package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 噶尔GarBarony struct {
	feud.BaseBarony
}

var BGar噶尔 feud.Barony = &噶尔GarBarony{}

func init() {
	f := BGar噶尔.(*噶尔GarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gar",
		TitleName: "噶尔",
		TitleCode: "b_gar",
	}
}
