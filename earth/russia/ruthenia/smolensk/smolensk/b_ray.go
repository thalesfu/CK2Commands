package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖RayBarony struct {
	feud.BaseBarony
}

var BRay赖 feud.Barony = &赖RayBarony{}

func init() {
    f := BRay赖.(*赖RayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ray",
		TitleName: "赖",
		TitleCode: "b_ray",
	}
}
