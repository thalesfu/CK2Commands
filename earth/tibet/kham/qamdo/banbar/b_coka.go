package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 草卡CokaBarony struct {
	feud.BaseBarony
}

var BCoka草卡 feud.Barony = &草卡CokaBarony{}

func init() {
	f := BCoka草卡.(*草卡CokaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coka",
		TitleName: "草卡",
		TitleCode: "b_coka",
	}
}
