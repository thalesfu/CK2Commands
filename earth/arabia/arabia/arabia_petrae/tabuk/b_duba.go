package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜巴DubaBarony struct {
	feud.BaseBarony
}

var BDuba杜巴 feud.Barony = &杜巴DubaBarony{}

func init() {
    f := BDuba杜巴.(*杜巴DubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duba",
		TitleName: "杜巴",
		TitleCode: "b_duba",
	}
}
