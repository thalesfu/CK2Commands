package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎西岗TrashigangBarony struct {
	feud.BaseBarony
}

var BTrashigang扎西岗 feud.Barony = &扎西岗TrashigangBarony{}

func init() {
    f := BTrashigang扎西岗.(*扎西岗TrashigangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trashigang",
		TitleName: "扎西岗",
		TitleCode: "b_trashigang",
	}
}
