package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波西塔诺PositanoBarony struct {
	feud.BaseBarony
}

var BPositano波西塔诺 feud.Barony = &波西塔诺PositanoBarony{}

func init() {
    f := BPositano波西塔诺.(*波西塔诺PositanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "positano",
		TitleName: "波西塔诺",
		TitleCode: "b_positano",
	}
}
