package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔代尔KildareBarony struct {
	feud.BaseBarony
}

var BKildare基尔代尔 feud.Barony = &基尔代尔KildareBarony{}

func init() {
	f := BKildare基尔代尔.(*基尔代尔KildareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kildare",
		TitleName: "基尔代尔",
		TitleCode: "b_kildare",
	}
}
