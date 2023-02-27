package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦卢奥纳VelluonaBarony struct {
	feud.BaseBarony
}

var BVelluona韦卢奥纳 feud.Barony = &韦卢奥纳VelluonaBarony{}

func init() {
    f := BVelluona韦卢奥纳.(*韦卢奥纳VelluonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velluona",
		TitleName: "韦卢奥纳",
		TitleCode: "b_velluona",
	}
}
