package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福廷格尔FortingallBarony struct {
	feud.BaseBarony
}

var BFortingall福廷格尔 feud.Barony = &福廷格尔FortingallBarony{}

func init() {
	f := BFortingall福廷格尔.(*福廷格尔FortingallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fortingall",
		TitleName: "福廷格尔",
		TitleCode: "b_fortingall",
	}
}
