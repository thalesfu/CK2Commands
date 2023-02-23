package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔伊夫TaifBarony struct {
	feud.BaseBarony
}

var BTaif塔伊夫 feud.Barony = &塔伊夫TaifBarony{}

func init() {
	f := BTaif塔伊夫.(*塔伊夫TaifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taif",
		TitleName: "塔伊夫",
		TitleCode: "b_taif",
	}
}
