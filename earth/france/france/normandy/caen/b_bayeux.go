package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴约BayeuxBarony struct {
	feud.BaseBarony
}

var BBayeux巴约 feud.Barony = &巴约BayeuxBarony{}

func init() {
	f := BBayeux巴约.(*巴约BayeuxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayeux",
		TitleName: "巴约",
		TitleCode: "b_bayeux",
	}
}
