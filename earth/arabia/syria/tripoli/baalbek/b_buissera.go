package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布瑟塞拉BuisseraBarony struct {
	feud.BaseBarony
}

var BBuissera布瑟塞拉 feud.Barony = &布瑟塞拉BuisseraBarony{}

func init() {
    f := BBuissera布瑟塞拉.(*布瑟塞拉BuisseraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buissera",
		TitleName: "布瑟塞拉",
		TitleCode: "b_buissera",
	}
}
