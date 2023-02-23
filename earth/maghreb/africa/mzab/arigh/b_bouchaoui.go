package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布沙维BouchaouiBarony struct {
	feud.BaseBarony
}

var BBouchaoui布沙维 feud.Barony = &布沙维BouchaouiBarony{}

func init() {
	f := BBouchaoui布沙维.(*布沙维BouchaouiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bouchaoui",
		TitleName: "布沙维",
		TitleCode: "b_bouchaoui",
	}
}
