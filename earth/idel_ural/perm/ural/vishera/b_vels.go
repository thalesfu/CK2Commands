package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔斯VelsBarony struct {
	feud.BaseBarony
}

var BVels韦尔斯 feud.Barony = &韦尔斯VelsBarony{}

func init() {
    f := BVels韦尔斯.(*韦尔斯VelsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vels",
		TitleName: "韦尔斯",
		TitleCode: "b_vels",
	}
}
