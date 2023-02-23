package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧村AultBarony struct {
	feud.BaseBarony
}

var BAult欧村 feud.Barony = &欧村AultBarony{}

func init() {
	f := BAult欧村.(*欧村AultBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ault",
		TitleName: "欧村",
		TitleCode: "b_ault",
	}
}
