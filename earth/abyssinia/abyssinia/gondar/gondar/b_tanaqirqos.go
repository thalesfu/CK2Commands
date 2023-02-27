package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔纳奇果斯TanaqirqosBarony struct {
	feud.BaseBarony
}

var BTanaqirqos塔纳奇果斯 feud.Barony = &塔纳奇果斯TanaqirqosBarony{}

func init() {
    f := BTanaqirqos塔纳奇果斯.(*塔纳奇果斯TanaqirqosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanaqirqos",
		TitleName: "塔纳奇果斯",
		TitleCode: "b_tanaqirqos",
	}
}
