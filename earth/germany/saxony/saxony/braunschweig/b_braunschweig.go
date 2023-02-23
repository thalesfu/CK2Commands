package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布伦瑞克BraunschweigBarony struct {
	feud.BaseBarony
}

var BBraunschweig布伦瑞克 feud.Barony = &布伦瑞克BraunschweigBarony{}

func init() {
	f := BBraunschweig布伦瑞克.(*布伦瑞克BraunschweigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "braunschweig",
		TitleName: "布伦瑞克",
		TitleCode: "b_braunschweig",
	}
}
