package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因塔IntaBarony struct {
	feud.BaseBarony
}

var BInta因塔 feud.Barony = &因塔IntaBarony{}

func init() {
    f := BInta因塔.(*因塔IntaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inta",
		TitleName: "因塔",
		TitleCode: "b_inta",
	}
}
