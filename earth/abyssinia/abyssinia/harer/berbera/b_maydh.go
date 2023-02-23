package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈德MaydhBarony struct {
	feud.BaseBarony
}

var BMaydh迈德 feud.Barony = &迈德MaydhBarony{}

func init() {
	f := BMaydh迈德.(*迈德MaydhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maydh",
		TitleName: "迈德",
		TitleCode: "b_maydh",
	}
}
