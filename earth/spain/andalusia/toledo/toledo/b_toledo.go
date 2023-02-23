package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托雷多ToledoBarony struct {
	feud.BaseBarony
}

var BToledo托雷多 feud.Barony = &托雷多ToledoBarony{}

func init() {
	f := BToledo托雷多.(*托雷多ToledoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toledo",
		TitleName: "托雷多",
		TitleCode: "b_toledo",
	}
}
