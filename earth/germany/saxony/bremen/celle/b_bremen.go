package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 不来梅BremenBarony struct {
	feud.BaseBarony
}

var BBremen不来梅 feud.Barony = &不来梅BremenBarony{}

func init() {
	f := BBremen不来梅.(*不来梅BremenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bremen",
		TitleName: "不来梅",
		TitleCode: "b_bremen",
	}
}
