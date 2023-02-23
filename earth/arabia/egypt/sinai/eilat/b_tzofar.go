package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琐法TzofarBarony struct {
	feud.BaseBarony
}

var BTzofar琐法 feud.Barony = &琐法TzofarBarony{}

func init() {
	f := BTzofar琐法.(*琐法TzofarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tzofar",
		TitleName: "琐法",
		TitleCode: "b_tzofar",
	}
}
