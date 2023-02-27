package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔切斯特ColchesterBarony struct {
	feud.BaseBarony
}

var BColchester科尔切斯特 feud.Barony = &科尔切斯特ColchesterBarony{}

func init() {
    f := BColchester科尔切斯特.(*科尔切斯特ColchesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colchester",
		TitleName: "科尔切斯特",
		TitleCode: "b_colchester",
	}
}
