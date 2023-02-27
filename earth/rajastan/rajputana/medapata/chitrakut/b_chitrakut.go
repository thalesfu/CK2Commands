package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 质多罗矩吒ChitrakutBarony struct {
	feud.BaseBarony
}

var BChitrakut质多罗矩吒 feud.Barony = &质多罗矩吒ChitrakutBarony{}

func init() {
    f := BChitrakut质多罗矩吒.(*质多罗矩吒ChitrakutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chitrakut",
		TitleName: "质多罗矩吒",
		TitleCode: "b_chitrakut",
	}
}
