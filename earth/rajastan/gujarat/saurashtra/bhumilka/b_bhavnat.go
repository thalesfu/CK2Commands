package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般那特BhavnatBarony struct {
	feud.BaseBarony
}

var BBhavnat般那特 feud.Barony = &般那特BhavnatBarony{}

func init() {
    f := BBhavnat般那特.(*般那特BhavnatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhavnat",
		TitleName: "般那特",
		TitleCode: "b_bhavnat",
	}
}
