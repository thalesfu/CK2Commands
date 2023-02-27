package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏夫SufBarony struct {
	feud.BaseBarony
}

var BSuf苏夫 feud.Barony = &苏夫SufBarony{}

func init() {
    f := BSuf苏夫.(*苏夫SufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suf",
		TitleName: "苏夫",
		TitleCode: "b_suf",
	}
}
