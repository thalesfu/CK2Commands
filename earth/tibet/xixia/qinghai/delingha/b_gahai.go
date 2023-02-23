package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尕海GahaiBarony struct {
	feud.BaseBarony
}

var BGahai尕海 feud.Barony = &尕海GahaiBarony{}

func init() {
	f := BGahai尕海.(*尕海GahaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gahai",
		TitleName: "尕海",
		TitleCode: "b_gahai",
	}
}
