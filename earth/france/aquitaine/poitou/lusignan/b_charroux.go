package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙鲁CharrouxBarony struct {
	feud.BaseBarony
}

var BCharroux沙鲁 feud.Barony = &沙鲁CharrouxBarony{}

func init() {
	f := BCharroux沙鲁.(*沙鲁CharrouxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charroux",
		TitleName: "沙鲁",
		TitleCode: "b_charroux",
	}
}
