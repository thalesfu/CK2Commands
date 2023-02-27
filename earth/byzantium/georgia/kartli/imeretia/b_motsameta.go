package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫察梅塔MotsametaBarony struct {
	feud.BaseBarony
}

var BMotsameta莫察梅塔 feud.Barony = &莫察梅塔MotsametaBarony{}

func init() {
    f := BMotsameta莫察梅塔.(*莫察梅塔MotsametaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "motsameta",
		TitleName: "莫察梅塔",
		TitleCode: "b_motsameta",
	}
}
