package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄克特UrquhartBarony struct {
	feud.BaseBarony
}

var BUrquhart厄克特 feud.Barony = &厄克特UrquhartBarony{}

func init() {
	f := BUrquhart厄克特.(*厄克特UrquhartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urquhart",
		TitleName: "厄克特",
		TitleCode: "b_urquhart",
	}
}
