package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉喀什MarrakechBarony struct {
	feud.BaseBarony
}

var BMarrakech马拉喀什 feud.Barony = &马拉喀什MarrakechBarony{}

func init() {
	f := BMarrakech马拉喀什.(*马拉喀什MarrakechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marrakech",
		TitleName: "马拉喀什",
		TitleCode: "b_marrakech",
	}
}
