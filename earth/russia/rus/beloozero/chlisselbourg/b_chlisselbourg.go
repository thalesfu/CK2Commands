package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什利谢利堡ChlisselbourgBarony struct {
	feud.BaseBarony
}

var BChlisselbourg什利谢利堡 feud.Barony = &什利谢利堡ChlisselbourgBarony{}

func init() {
	f := BChlisselbourg什利谢利堡.(*什利谢利堡ChlisselbourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chlisselbourg",
		TitleName: "什利谢利堡",
		TitleCode: "b_chlisselbourg",
	}
}
