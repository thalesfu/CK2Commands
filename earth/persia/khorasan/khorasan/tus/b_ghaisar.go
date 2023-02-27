package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖萨尔GhaisarBarony struct {
	feud.BaseBarony
}

var BGhaisar盖萨尔 feud.Barony = &盖萨尔GhaisarBarony{}

func init() {
    f := BGhaisar盖萨尔.(*盖萨尔GhaisarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghaisar",
		TitleName: "盖萨尔",
		TitleCode: "b_ghaisar",
	}
}
