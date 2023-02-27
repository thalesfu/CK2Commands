package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尼苏Guunii_suuBarony struct {
	feud.BaseBarony
}

var BGuunii_suu古尼苏 feud.Barony = &古尼苏Guunii_suuBarony{}

func init() {
    f := BGuunii_suu古尼苏.(*古尼苏Guunii_suuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guunii_suu",
		TitleName: "古尼苏",
		TitleCode: "b_guunii_suu",
	}
}
