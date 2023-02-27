package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布瓦南BouananeBarony struct {
	feud.BaseBarony
}

var BBouanane布瓦南 feud.Barony = &布瓦南BouananeBarony{}

func init() {
    f := BBouanane布瓦南.(*布瓦南BouananeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bouanane",
		TitleName: "布瓦南",
		TitleCode: "b_bouanane",
	}
}
