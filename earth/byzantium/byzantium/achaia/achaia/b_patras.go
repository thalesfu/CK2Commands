package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕特拉斯PatrasBarony struct {
	feud.BaseBarony
}

var BPatras帕特拉斯 feud.Barony = &帕特拉斯PatrasBarony{}

func init() {
	f := BPatras帕特拉斯.(*帕特拉斯PatrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patras",
		TitleName: "帕特拉斯",
		TitleCode: "b_patras",
	}
}
