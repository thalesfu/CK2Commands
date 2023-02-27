package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马堡Marburg_pettauBarony struct {
	feud.BaseBarony
}

var BMarburg_pettau马堡 feud.Barony = &马堡Marburg_pettauBarony{}

func init() {
    f := BMarburg_pettau马堡.(*马堡Marburg_pettauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marburg_pettau",
		TitleName: "马堡",
		TitleCode: "b_marburg_pettau",
	}
}
