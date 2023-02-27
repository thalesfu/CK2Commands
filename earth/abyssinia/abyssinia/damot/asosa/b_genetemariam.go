package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 根特迈利GenetemariamBarony struct {
	feud.BaseBarony
}

var BGenetemariam根特迈利 feud.Barony = &根特迈利GenetemariamBarony{}

func init() {
    f := BGenetemariam根特迈利.(*根特迈利GenetemariamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "genetemariam",
		TitleName: "根特迈利",
		TitleCode: "b_genetemariam",
	}
}
