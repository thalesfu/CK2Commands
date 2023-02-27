package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃毗犍阇HabiganjBarony struct {
	feud.BaseBarony
}

var BHabiganj诃毗犍阇 feud.Barony = &诃毗犍阇HabiganjBarony{}

func init() {
    f := BHabiganj诃毗犍阇.(*诃毗犍阇HabiganjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "habiganj",
		TitleName: "诃毗犍阇",
		TitleCode: "b_habiganj",
	}
}
