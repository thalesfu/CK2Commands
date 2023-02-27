package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加埃塔GaetaBarony struct {
	feud.BaseBarony
}

var BGaeta加埃塔 feud.Barony = &加埃塔GaetaBarony{}

func init() {
    f := BGaeta加埃塔.(*加埃塔GaetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaeta",
		TitleName: "加埃塔",
		TitleCode: "b_gaeta",
	}
}
