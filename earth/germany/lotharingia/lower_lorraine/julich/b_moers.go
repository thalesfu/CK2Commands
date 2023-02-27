package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默尔斯MoersBarony struct {
	feud.BaseBarony
}

var BMoers默尔斯 feud.Barony = &默尔斯MoersBarony{}

func init() {
    f := BMoers默尔斯.(*默尔斯MoersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moers",
		TitleName: "默尔斯",
		TitleCode: "b_moers",
	}
}
