package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奎尔富特QuerfurtBarony struct {
	feud.BaseBarony
}

var BQuerfurt奎尔富特 feud.Barony = &奎尔富特QuerfurtBarony{}

func init() {
	f := BQuerfurt奎尔富特.(*奎尔富特QuerfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "querfurt",
		TitleName: "奎尔富特",
		TitleCode: "b_querfurt",
	}
}
