package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利布尔讷LibourneBarony struct {
	feud.BaseBarony
}

var BLibourne利布尔讷 feud.Barony = &利布尔讷LibourneBarony{}

func init() {
	f := BLibourne利布尔讷.(*利布尔讷LibourneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "libourne",
		TitleName: "利布尔讷",
		TitleCode: "b_libourne",
	}
}
