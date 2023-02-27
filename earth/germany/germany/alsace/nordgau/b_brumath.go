package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布吕马特BrumathBarony struct {
	feud.BaseBarony
}

var BBrumath布吕马特 feud.Barony = &布吕马特BrumathBarony{}

func init() {
    f := BBrumath布吕马特.(*布吕马特BrumathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brumath",
		TitleName: "布吕马特",
		TitleCode: "b_brumath",
	}
}
