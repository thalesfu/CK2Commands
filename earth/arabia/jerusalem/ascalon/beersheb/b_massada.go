package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨达MassadaBarony struct {
	feud.BaseBarony
}

var BMassada马萨达 feud.Barony = &马萨达MassadaBarony{}

func init() {
	f := BMassada马萨达.(*马萨达MassadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "massada",
		TitleName: "马萨达",
		TitleCode: "b_massada",
	}
}
