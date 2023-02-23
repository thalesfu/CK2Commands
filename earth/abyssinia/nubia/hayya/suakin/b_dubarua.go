package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜巴鲁阿DubaruaBarony struct {
	feud.BaseBarony
}

var BDubarua杜巴鲁阿 feud.Barony = &杜巴鲁阿DubaruaBarony{}

func init() {
	f := BDubarua杜巴鲁阿.(*杜巴鲁阿DubaruaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubarua",
		TitleName: "杜巴鲁阿",
		TitleCode: "b_dubarua",
	}
}
