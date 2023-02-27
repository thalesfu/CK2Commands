package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔索纳SolsonaBarony struct {
	feud.BaseBarony
}

var BSolsona索尔索纳 feud.Barony = &索尔索纳SolsonaBarony{}

func init() {
    f := BSolsona索尔索纳.(*索尔索纳SolsonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solsona",
		TitleName: "索尔索纳",
		TitleCode: "b_solsona",
	}
}
