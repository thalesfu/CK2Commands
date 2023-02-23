package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍费尔HoffelBarony struct {
	feud.BaseBarony
}

var BHoffel霍费尔 feud.Barony = &霍费尔HoffelBarony{}

func init() {
	f := BHoffel霍费尔.(*霍费尔HoffelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hoffel",
		TitleName: "霍费尔",
		TitleCode: "b_hoffel",
	}
}
