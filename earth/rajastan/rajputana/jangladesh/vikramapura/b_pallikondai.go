package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波梨军荼PallikondaiBarony struct {
	feud.BaseBarony
}

var BPallikondai波梨军荼 feud.Barony = &波梨军荼PallikondaiBarony{}

func init() {
    f := BPallikondai波梨军荼.(*波梨军荼PallikondaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pallikondai",
		TitleName: "波梨军荼",
		TitleCode: "b_pallikondai",
	}
}
