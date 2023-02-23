package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒门ShumenBarony struct {
	feud.BaseBarony
}

var BShumen舒门 feud.Barony = &舒门ShumenBarony{}

func init() {
	f := BShumen舒门.(*舒门ShumenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shumen",
		TitleName: "舒门",
		TitleCode: "b_shumen",
	}
}
