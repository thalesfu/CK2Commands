package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赛MarseilleBarony struct {
	feud.BaseBarony
}

var BMarseille马赛 feud.Barony = &马赛MarseilleBarony{}

func init() {
	f := BMarseille马赛.(*马赛MarseilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marseille",
		TitleName: "马赛",
		TitleCode: "b_marseille",
	}
}
