package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩迦罗MagharBarony struct {
	feud.BaseBarony
}

var BMaghar摩迦罗 feud.Barony = &摩迦罗MagharBarony{}

func init() {
    f := BMaghar摩迦罗.(*摩迦罗MagharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maghar",
		TitleName: "摩迦罗",
		TitleCode: "b_maghar",
	}
}
