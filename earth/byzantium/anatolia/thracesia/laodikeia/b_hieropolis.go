package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑洛波利斯HieropolisBarony struct {
	feud.BaseBarony
}

var BHieropolis黑洛波利斯 feud.Barony = &黑洛波利斯HieropolisBarony{}

func init() {
	f := BHieropolis黑洛波利斯.(*黑洛波利斯HieropolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hieropolis",
		TitleName: "黑洛波利斯",
		TitleCode: "b_hieropolis",
	}
}
