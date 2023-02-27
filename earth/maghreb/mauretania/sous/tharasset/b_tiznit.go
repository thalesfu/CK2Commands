package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提兹尼特TiznitBarony struct {
	feud.BaseBarony
}

var BTiznit提兹尼特 feud.Barony = &提兹尼特TiznitBarony{}

func init() {
    f := BTiznit提兹尼特.(*提兹尼特TiznitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiznit",
		TitleName: "提兹尼特",
		TitleCode: "b_tiznit",
	}
}
