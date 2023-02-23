package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多德霍特DoderhultBarony struct {
	feud.BaseBarony
}

var BDoderhult多德霍特 feud.Barony = &多德霍特DoderhultBarony{}

func init() {
	f := BDoderhult多德霍特.(*多德霍特DoderhultBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doderhult",
		TitleName: "多德霍特",
		TitleCode: "b_doderhult",
	}
}
