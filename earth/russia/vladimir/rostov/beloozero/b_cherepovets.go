package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切列波韦茨CherepovetsBarony struct {
	feud.BaseBarony
}

var BCherepovets切列波韦茨 feud.Barony = &切列波韦茨CherepovetsBarony{}

func init() {
	f := BCherepovets切列波韦茨.(*切列波韦茨CherepovetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherepovets",
		TitleName: "切列波韦茨",
		TitleCode: "b_cherepovets",
	}
}
