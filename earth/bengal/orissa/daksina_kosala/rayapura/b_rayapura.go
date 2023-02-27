package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗耶补罗RayapuraBarony struct {
	feud.BaseBarony
}

var BRayapura罗耶补罗 feud.Barony = &罗耶补罗RayapuraBarony{}

func init() {
    f := BRayapura罗耶补罗.(*罗耶补罗RayapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rayapura",
		TitleName: "罗耶补罗",
		TitleCode: "b_rayapura",
	}
}
