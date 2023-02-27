package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑金SanginBarony struct {
	feud.BaseBarony
}

var BSangin桑金 feud.Barony = &桑金SanginBarony{}

func init() {
    f := BSangin桑金.(*桑金SanginBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangin",
		TitleName: "桑金",
		TitleCode: "b_sangin",
	}
}
