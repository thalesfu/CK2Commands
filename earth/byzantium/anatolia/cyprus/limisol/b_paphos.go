package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕福斯PaphosBarony struct {
	feud.BaseBarony
}

var BPaphos帕福斯 feud.Barony = &帕福斯PaphosBarony{}

func init() {
	f := BPaphos帕福斯.(*帕福斯PaphosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paphos",
		TitleName: "帕福斯",
		TitleCode: "b_paphos",
	}
}
