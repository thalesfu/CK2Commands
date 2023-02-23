package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔德MoldeBarony struct {
	feud.BaseBarony
}

var BMolde莫尔德 feud.Barony = &莫尔德MoldeBarony{}

func init() {
	f := BMolde莫尔德.(*莫尔德MoldeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molde",
		TitleName: "莫尔德",
		TitleCode: "b_molde",
	}
}
