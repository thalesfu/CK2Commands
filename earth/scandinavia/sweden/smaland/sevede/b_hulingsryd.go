package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡令斯吕德HulingsrydBarony struct {
	feud.BaseBarony
}

var BHulingsryd胡令斯吕德 feud.Barony = &胡令斯吕德HulingsrydBarony{}

func init() {
	f := BHulingsryd胡令斯吕德.(*胡令斯吕德HulingsrydBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hulingsryd",
		TitleName: "胡令斯吕德",
		TitleCode: "b_hulingsryd",
	}
}
