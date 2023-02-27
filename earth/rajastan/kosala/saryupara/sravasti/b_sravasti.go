package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍卫城SravastiBarony struct {
	feud.BaseBarony
}

var BSravasti舍卫城 feud.Barony = &舍卫城SravastiBarony{}

func init() {
    f := BSravasti舍卫城.(*舍卫城SravastiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sravasti",
		TitleName: "舍卫城",
		TitleCode: "b_sravasti",
	}
}
