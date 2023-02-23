package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利林费尔德LilienfeldBarony struct {
	feud.BaseBarony
}

var BLilienfeld利林费尔德 feud.Barony = &利林费尔德LilienfeldBarony{}

func init() {
	f := BLilienfeld利林费尔德.(*利林费尔德LilienfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lilienfeld",
		TitleName: "利林费尔德",
		TitleCode: "b_lilienfeld",
	}
}
