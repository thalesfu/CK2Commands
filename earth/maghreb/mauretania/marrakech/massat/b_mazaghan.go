package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马扎汗MazaghanBarony struct {
	feud.BaseBarony
}

var BMazaghan马扎汗 feud.Barony = &马扎汗MazaghanBarony{}

func init() {
    f := BMazaghan马扎汗.(*马扎汗MazaghanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazaghan",
		TitleName: "马扎汗",
		TitleCode: "b_mazaghan",
	}
}
