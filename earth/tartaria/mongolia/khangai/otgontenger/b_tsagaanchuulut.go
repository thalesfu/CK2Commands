package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察干楚卢特TsagaanchuulutBarony struct {
	feud.BaseBarony
}

var BTsagaanchuulut察干楚卢特 feud.Barony = &察干楚卢特TsagaanchuulutBarony{}

func init() {
	f := BTsagaanchuulut察干楚卢特.(*察干楚卢特TsagaanchuulutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsagaanchuulut",
		TitleName: "察干楚卢特",
		TitleCode: "b_tsagaanchuulut",
	}
}
