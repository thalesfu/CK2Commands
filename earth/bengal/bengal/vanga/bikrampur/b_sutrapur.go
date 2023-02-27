package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 修多罗补罗SutrapurBarony struct {
	feud.BaseBarony
}

var BSutrapur修多罗补罗 feud.Barony = &修多罗补罗SutrapurBarony{}

func init() {
    f := BSutrapur修多罗补罗.(*修多罗补罗SutrapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sutrapur",
		TitleName: "修多罗补罗",
		TitleCode: "b_sutrapur",
	}
}
