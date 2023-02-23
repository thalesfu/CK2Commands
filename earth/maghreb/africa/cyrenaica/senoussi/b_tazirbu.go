package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔济尔布TazirbuBarony struct {
	feud.BaseBarony
}

var BTazirbu塔济尔布 feud.Barony = &塔济尔布TazirbuBarony{}

func init() {
	f := BTazirbu塔济尔布.(*塔济尔布TazirbuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tazirbu",
		TitleName: "塔济尔布",
		TitleCode: "b_tazirbu",
	}
}
