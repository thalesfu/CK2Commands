package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安古洛波利斯ArgyropolisBarony struct {
	feud.BaseBarony
}

var BArgyropolis安古洛波利斯 feud.Barony = &安古洛波利斯ArgyropolisBarony{}

func init() {
	f := BArgyropolis安古洛波利斯.(*安古洛波利斯ArgyropolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argyropolis",
		TitleName: "安古洛波利斯",
		TitleCode: "b_argyropolis",
	}
}
