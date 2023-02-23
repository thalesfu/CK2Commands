package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺勒姆NorhamBarony struct {
	feud.BaseBarony
}

var BNorham诺勒姆 feud.Barony = &诺勒姆NorhamBarony{}

func init() {
	f := BNorham诺勒姆.(*诺勒姆NorhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norham",
		TitleName: "诺勒姆",
		TitleCode: "b_norham",
	}
}
