package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马因MainBarony struct {
	feud.BaseBarony
}

var BMain马因 feud.Barony = &马因MainBarony{}

func init() {
	f := BMain马因.(*马因MainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "main",
		TitleName: "马因",
		TitleCode: "b_main",
	}
}
