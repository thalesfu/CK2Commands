package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫萨利斯克MosalskBarony struct {
	feud.BaseBarony
}

var BMosalsk莫萨利斯克 feud.Barony = &莫萨利斯克MosalskBarony{}

func init() {
	f := BMosalsk莫萨利斯克.(*莫萨利斯克MosalskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosalsk",
		TitleName: "莫萨利斯克",
		TitleCode: "b_mosalsk",
	}
}
