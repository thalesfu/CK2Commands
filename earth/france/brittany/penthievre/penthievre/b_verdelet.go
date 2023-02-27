package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔德莱VerdeletBarony struct {
	feud.BaseBarony
}

var BVerdelet韦尔德莱 feud.Barony = &韦尔德莱VerdeletBarony{}

func init() {
    f := BVerdelet韦尔德莱.(*韦尔德莱VerdeletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verdelet",
		TitleName: "韦尔德莱",
		TitleCode: "b_verdelet",
	}
}
