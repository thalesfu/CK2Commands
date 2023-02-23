package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘塔斯GantasBarony struct {
	feud.BaseBarony
}

var BGantas甘塔斯 feud.Barony = &甘塔斯GantasBarony{}

func init() {
	f := BGantas甘塔斯.(*甘塔斯GantasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gantas",
		TitleName: "甘塔斯",
		TitleCode: "b_gantas",
	}
}
