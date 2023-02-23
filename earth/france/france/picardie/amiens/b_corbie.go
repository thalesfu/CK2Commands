package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔比CorbieBarony struct {
	feud.BaseBarony
}

var BCorbie科尔比 feud.Barony = &科尔比CorbieBarony{}

func init() {
	f := BCorbie科尔比.(*科尔比CorbieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corbie",
		TitleName: "科尔比",
		TitleCode: "b_corbie",
	}
}
