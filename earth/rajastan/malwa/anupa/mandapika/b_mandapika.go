package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼荼卑迦MandapikaBarony struct {
	feud.BaseBarony
}

var BMandapika曼荼卑迦 feud.Barony = &曼荼卑迦MandapikaBarony{}

func init() {
	f := BMandapika曼荼卑迦.(*曼荼卑迦MandapikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandapika",
		TitleName: "曼荼卑迦",
		TitleCode: "b_mandapika",
	}
}
