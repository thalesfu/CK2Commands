package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔布OrbeBarony struct {
	feud.BaseBarony
}

var BOrbe奥尔布 feud.Barony = &奥尔布OrbeBarony{}

func init() {
    f := BOrbe奥尔布.(*奥尔布OrbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orbe",
		TitleName: "奥尔布",
		TitleCode: "b_orbe",
	}
}
