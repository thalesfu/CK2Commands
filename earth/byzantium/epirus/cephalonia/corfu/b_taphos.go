package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔褔斯TaphosBarony struct {
	feud.BaseBarony
}

var BTaphos塔褔斯 feud.Barony = &塔褔斯TaphosBarony{}

func init() {
	f := BTaphos塔褔斯.(*塔褔斯TaphosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taphos",
		TitleName: "塔褔斯",
		TitleCode: "b_taphos",
	}
}
