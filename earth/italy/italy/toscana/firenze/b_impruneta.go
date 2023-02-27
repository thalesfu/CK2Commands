package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因普鲁内塔ImprunetaBarony struct {
	feud.BaseBarony
}

var BImpruneta因普鲁内塔 feud.Barony = &因普鲁内塔ImprunetaBarony{}

func init() {
    f := BImpruneta因普鲁内塔.(*因普鲁内塔ImprunetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "impruneta",
		TitleName: "因普鲁内塔",
		TitleCode: "b_impruneta",
	}
}
