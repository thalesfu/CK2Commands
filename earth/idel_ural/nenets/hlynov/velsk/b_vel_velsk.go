package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔Vel_velskBarony struct {
	feud.BaseBarony
}

var BVel_velsk韦尔 feud.Barony = &韦尔Vel_velskBarony{}

func init() {
    f := BVel_velsk韦尔.(*韦尔Vel_velskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vel_velsk",
		TitleName: "韦尔",
		TitleCode: "b_vel_velsk",
	}
}
