package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊维萨EivissaBarony struct {
	feud.BaseBarony
}

var BEivissa伊维萨 feud.Barony = &伊维萨EivissaBarony{}

func init() {
	f := BEivissa伊维萨.(*伊维萨EivissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eivissa",
		TitleName: "伊维萨",
		TitleCode: "b_eivissa",
	}
}
