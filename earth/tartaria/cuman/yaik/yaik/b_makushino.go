package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马库希诺MakushinoBarony struct {
	feud.BaseBarony
}

var BMakushino马库希诺 feud.Barony = &马库希诺MakushinoBarony{}

func init() {
	f := BMakushino马库希诺.(*马库希诺MakushinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makushino",
		TitleName: "马库希诺",
		TitleCode: "b_makushino",
	}
}
