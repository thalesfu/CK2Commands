package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴多戈BadogoBarony struct {
	feud.BaseBarony
}

var BBadogo巴多戈 feud.Barony = &巴多戈BadogoBarony{}

func init() {
	f := BBadogo巴多戈.(*巴多戈BadogoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badogo",
		TitleName: "巴多戈",
		TitleCode: "b_badogo",
	}
}
