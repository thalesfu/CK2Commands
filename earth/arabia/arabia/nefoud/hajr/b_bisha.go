package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比沙BishaBarony struct {
	feud.BaseBarony
}

var BBisha比沙 feud.Barony = &比沙BishaBarony{}

func init() {
    f := BBisha比沙.(*比沙BishaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bisha",
		TitleName: "比沙",
		TitleCode: "b_bisha",
	}
}
