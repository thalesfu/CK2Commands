package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛迦诺LocarnoBarony struct {
	feud.BaseBarony
}

var BLocarno洛迦诺 feud.Barony = &洛迦诺LocarnoBarony{}

func init() {
	f := BLocarno洛迦诺.(*洛迦诺LocarnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "locarno",
		TitleName: "洛迦诺",
		TitleCode: "b_locarno",
	}
}
