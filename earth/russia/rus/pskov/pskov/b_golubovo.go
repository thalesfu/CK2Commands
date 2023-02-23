package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈卢博沃GolubovoBarony struct {
	feud.BaseBarony
}

var BGolubovo戈卢博沃 feud.Barony = &戈卢博沃GolubovoBarony{}

func init() {
	f := BGolubovo戈卢博沃.(*戈卢博沃GolubovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golubovo",
		TitleName: "戈卢博沃",
		TitleCode: "b_golubovo",
	}
}
