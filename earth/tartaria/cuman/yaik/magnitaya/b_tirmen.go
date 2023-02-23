package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季尔缅TirmenBarony struct {
	feud.BaseBarony
}

var BTirmen季尔缅 feud.Barony = &季尔缅TirmenBarony{}

func init() {
	f := BTirmen季尔缅.(*季尔缅TirmenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirmen",
		TitleName: "季尔缅",
		TitleCode: "b_tirmen",
	}
}
