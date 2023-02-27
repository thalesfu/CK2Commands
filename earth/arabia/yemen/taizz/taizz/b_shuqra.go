package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒拉ShuqraBarony struct {
	feud.BaseBarony
}

var BShuqra舒拉 feud.Barony = &舒拉ShuqraBarony{}

func init() {
    f := BShuqra舒拉.(*舒拉ShuqraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shuqra",
		TitleName: "舒拉",
		TitleCode: "b_shuqra",
	}
}
