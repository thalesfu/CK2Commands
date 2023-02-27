package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦莱赫拉德VelehradBarony struct {
	feud.BaseBarony
}

var BVelehrad韦莱赫拉德 feud.Barony = &韦莱赫拉德VelehradBarony{}

func init() {
    f := BVelehrad韦莱赫拉德.(*韦莱赫拉德VelehradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velehrad",
		TitleName: "韦莱赫拉德",
		TitleCode: "b_velehrad",
	}
}
