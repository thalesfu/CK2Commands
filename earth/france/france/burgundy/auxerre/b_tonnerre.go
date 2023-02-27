package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托内尔TonnerreBarony struct {
	feud.BaseBarony
}

var BTonnerre托内尔 feud.Barony = &托内尔TonnerreBarony{}

func init() {
    f := BTonnerre托内尔.(*托内尔TonnerreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonnerre",
		TitleName: "托内尔",
		TitleCode: "b_tonnerre",
	}
}
