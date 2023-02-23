package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奎师那揭罗KrishnagarBarony struct {
	feud.BaseBarony
}

var BKrishnagar奎师那揭罗 feud.Barony = &奎师那揭罗KrishnagarBarony{}

func init() {
	f := BKrishnagar奎师那揭罗.(*奎师那揭罗KrishnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krishnagar",
		TitleName: "奎师那揭罗",
		TitleCode: "b_krishnagar",
	}
}
