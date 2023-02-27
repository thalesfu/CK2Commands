package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕勒涅PalleneBarony struct {
	feud.BaseBarony
}

var BPallene帕勒涅 feud.Barony = &帕勒涅PalleneBarony{}

func init() {
    f := BPallene帕勒涅.(*帕勒涅PalleneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pallene",
		TitleName: "帕勒涅",
		TitleCode: "b_pallene",
	}
}
