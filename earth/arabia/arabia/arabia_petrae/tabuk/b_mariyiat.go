package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马里伊特MariyiatBarony struct {
	feud.BaseBarony
}

var BMariyiat马里伊特 feud.Barony = &马里伊特MariyiatBarony{}

func init() {
    f := BMariyiat马里伊特.(*马里伊特MariyiatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mariyiat",
		TitleName: "马里伊特",
		TitleCode: "b_mariyiat",
	}
}
