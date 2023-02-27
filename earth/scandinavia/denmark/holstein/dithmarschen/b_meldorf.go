package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔多夫MeldorfBarony struct {
	feud.BaseBarony
}

var BMeldorf梅尔多夫 feud.Barony = &梅尔多夫MeldorfBarony{}

func init() {
    f := BMeldorf梅尔多夫.(*梅尔多夫MeldorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meldorf",
		TitleName: "梅尔多夫",
		TitleCode: "b_meldorf",
	}
}
