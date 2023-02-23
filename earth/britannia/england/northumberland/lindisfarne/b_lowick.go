package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛伊克LowickBarony struct {
	feud.BaseBarony
}

var BLowick洛伊克 feud.Barony = &洛伊克LowickBarony{}

func init() {
	f := BLowick洛伊克.(*洛伊克LowickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lowick",
		TitleName: "洛伊克",
		TitleCode: "b_lowick",
	}
}
