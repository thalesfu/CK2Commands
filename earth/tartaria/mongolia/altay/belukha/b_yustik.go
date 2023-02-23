package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤斯季克YustikBarony struct {
	feud.BaseBarony
}

var BYustik尤斯季克 feud.Barony = &尤斯季克YustikBarony{}

func init() {
	f := BYustik尤斯季克.(*尤斯季克YustikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yustik",
		TitleName: "尤斯季克",
		TitleCode: "b_yustik",
	}
}
