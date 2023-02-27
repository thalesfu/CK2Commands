package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布格图卜BougtobBarony struct {
	feud.BaseBarony
}

var BBougtob布格图卜 feud.Barony = &布格图卜BougtobBarony{}

func init() {
    f := BBougtob布格图卜.(*布格图卜BougtobBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bougtob",
		TitleName: "布格图卜",
		TitleCode: "b_bougtob",
	}
}
