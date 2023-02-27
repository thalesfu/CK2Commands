package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塞塔CasertaBarony struct {
	feud.BaseBarony
}

var BCaserta卡塞塔 feud.Barony = &卡塞塔CasertaBarony{}

func init() {
    f := BCaserta卡塞塔.(*卡塞塔CasertaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caserta",
		TitleName: "卡塞塔",
		TitleCode: "b_caserta",
	}
}
