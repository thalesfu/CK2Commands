package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞伊达SaidaBarony struct {
	feud.BaseBarony
}

var BSaida塞伊达 feud.Barony = &塞伊达SaidaBarony{}

func init() {
    f := BSaida塞伊达.(*塞伊达SaidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saida",
		TitleName: "塞伊达",
		TitleCode: "b_saida",
	}
}
