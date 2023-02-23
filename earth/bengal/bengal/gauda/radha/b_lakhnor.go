package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 落诃奴罗LakhnorBarony struct {
	feud.BaseBarony
}

var BLakhnor落诃奴罗 feud.Barony = &落诃奴罗LakhnorBarony{}

func init() {
	f := BLakhnor落诃奴罗.(*落诃奴罗LakhnorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakhnor",
		TitleName: "落诃奴罗",
		TitleCode: "b_lakhnor",
	}
}
