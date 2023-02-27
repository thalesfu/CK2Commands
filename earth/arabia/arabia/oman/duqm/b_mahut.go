package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马胡尔MahutBarony struct {
	feud.BaseBarony
}

var BMahut马胡尔 feud.Barony = &马胡尔MahutBarony{}

func init() {
    f := BMahut马胡尔.(*马胡尔MahutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahut",
		TitleName: "马胡尔",
		TitleCode: "b_mahut",
	}
}
