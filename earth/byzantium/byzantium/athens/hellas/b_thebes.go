package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底比斯ThebesBarony struct {
	feud.BaseBarony
}

var BThebes底比斯 feud.Barony = &底比斯ThebesBarony{}

func init() {
	f := BThebes底比斯.(*底比斯ThebesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thebes",
		TitleName: "底比斯",
		TitleCode: "b_thebes",
	}
}
