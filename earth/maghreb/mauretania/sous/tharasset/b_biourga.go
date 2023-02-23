package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比乌格拉BiourgaBarony struct {
	feud.BaseBarony
}

var BBiourga比乌格拉 feud.Barony = &比乌格拉BiourgaBarony{}

func init() {
	f := BBiourga比乌格拉.(*比乌格拉BiourgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biourga",
		TitleName: "比乌格拉",
		TitleCode: "b_biourga",
	}
}
