package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔库塔VorkutaBarony struct {
	feud.BaseBarony
}

var BVorkuta沃尔库塔 feud.Barony = &沃尔库塔VorkutaBarony{}

func init() {
    f := BVorkuta沃尔库塔.(*沃尔库塔VorkutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vorkuta",
		TitleName: "沃尔库塔",
		TitleCode: "b_vorkuta",
	}
}
