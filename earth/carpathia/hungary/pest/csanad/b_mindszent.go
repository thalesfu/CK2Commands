package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明曾特MindszentBarony struct {
	feud.BaseBarony
}

var BMindszent明曾特 feud.Barony = &明曾特MindszentBarony{}

func init() {
	f := BMindszent明曾特.(*明曾特MindszentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mindszent",
		TitleName: "明曾特",
		TitleCode: "b_mindszent",
	}
}
