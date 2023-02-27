package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 质多罗矩吒ChitrakutaBarony struct {
	feud.BaseBarony
}

var BChitrakuta质多罗矩吒 feud.Barony = &质多罗矩吒ChitrakutaBarony{}

func init() {
    f := BChitrakuta质多罗矩吒.(*质多罗矩吒ChitrakutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chitrakuta",
		TitleName: "质多罗矩吒",
		TitleCode: "b_chitrakuta",
	}
}
