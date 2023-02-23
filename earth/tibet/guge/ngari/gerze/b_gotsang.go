package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古昌GotsangBarony struct {
	feud.BaseBarony
}

var BGotsang古昌 feud.Barony = &古昌GotsangBarony{}

func init() {
	f := BGotsang古昌.(*古昌GotsangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gotsang",
		TitleName: "古昌",
		TitleCode: "b_gotsang",
	}
}
