package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔VelBarony struct {
	feud.BaseBarony
}

var BVel韦尔 feud.Barony = &韦尔VelBarony{}

func init() {
    f := BVel韦尔.(*韦尔VelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vel",
		TitleName: "韦尔",
		TitleCode: "b_vel",
	}
}
