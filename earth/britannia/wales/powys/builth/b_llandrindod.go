package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰德林多德LlandrindodBarony struct {
	feud.BaseBarony
}

var BLlandrindod兰德林多德 feud.Barony = &兰德林多德LlandrindodBarony{}

func init() {
    f := BLlandrindod兰德林多德.(*兰德林多德LlandrindodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llandrindod",
		TitleName: "兰德林多德",
		TitleCode: "b_llandrindod",
	}
}
