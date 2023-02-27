package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆克登MukdenBarony struct {
	feud.BaseBarony
}

var BMukden穆克登 feud.Barony = &穆克登MukdenBarony{}

func init() {
    f := BMukden穆克登.(*穆克登MukdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mukden",
		TitleName: "穆克登",
		TitleCode: "b_mukden",
	}
}
