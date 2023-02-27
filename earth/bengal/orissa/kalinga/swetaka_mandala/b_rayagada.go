package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗耶迦陀RayagadaBarony struct {
	feud.BaseBarony
}

var BRayagada罗耶迦陀 feud.Barony = &罗耶迦陀RayagadaBarony{}

func init() {
    f := BRayagada罗耶迦陀.(*罗耶迦陀RayagadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rayagada",
		TitleName: "罗耶迦陀",
		TitleCode: "b_rayagada",
	}
}
