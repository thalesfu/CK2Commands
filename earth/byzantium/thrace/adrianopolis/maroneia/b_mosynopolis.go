package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫西诺波利斯MosynopolisBarony struct {
	feud.BaseBarony
}

var BMosynopolis莫西诺波利斯 feud.Barony = &莫西诺波利斯MosynopolisBarony{}

func init() {
	f := BMosynopolis莫西诺波利斯.(*莫西诺波利斯MosynopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosynopolis",
		TitleName: "莫西诺波利斯",
		TitleCode: "b_mosynopolis",
	}
}
