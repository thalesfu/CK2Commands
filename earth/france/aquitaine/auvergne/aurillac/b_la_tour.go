package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉图尔La_tourBarony struct {
	feud.BaseBarony
}

var BLa_tour拉图尔 feud.Barony = &拉图尔La_tourBarony{}

func init() {
    f := BLa_tour拉图尔.(*拉图尔La_tourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "la_tour",
		TitleName: "拉图尔",
		TitleCode: "b_la_tour",
	}
}
