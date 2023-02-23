package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 札布让TsaparangBarony struct {
	feud.BaseBarony
}

var BTsaparang札布让 feud.Barony = &札布让TsaparangBarony{}

func init() {
	f := BTsaparang札布让.(*札布让TsaparangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsaparang",
		TitleName: "札布让",
		TitleCode: "b_tsaparang",
	}
}
