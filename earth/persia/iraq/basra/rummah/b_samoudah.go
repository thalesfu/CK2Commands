package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨穆达SamoudahBarony struct {
	feud.BaseBarony
}

var BSamoudah萨穆达 feud.Barony = &萨穆达SamoudahBarony{}

func init() {
	f := BSamoudah萨穆达.(*萨穆达SamoudahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samoudah",
		TitleName: "萨穆达",
		TitleCode: "b_samoudah",
	}
}
