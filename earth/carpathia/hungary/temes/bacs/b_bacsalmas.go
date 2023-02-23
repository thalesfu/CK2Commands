package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴乔尔马什BacsalmasBarony struct {
	feud.BaseBarony
}

var BBacsalmas巴乔尔马什 feud.Barony = &巴乔尔马什BacsalmasBarony{}

func init() {
	f := BBacsalmas巴乔尔马什.(*巴乔尔马什BacsalmasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bacsalmas",
		TitleName: "巴乔尔马什",
		TitleCode: "b_bacsalmas",
	}
}
