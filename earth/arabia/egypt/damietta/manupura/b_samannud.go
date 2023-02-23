package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨曼努德SamannudBarony struct {
	feud.BaseBarony
}

var BSamannud萨曼努德 feud.Barony = &萨曼努德SamannudBarony{}

func init() {
	f := BSamannud萨曼努德.(*萨曼努德SamannudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samannud",
		TitleName: "萨曼努德",
		TitleCode: "b_samannud",
	}
}
