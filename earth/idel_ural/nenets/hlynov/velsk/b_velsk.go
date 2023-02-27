package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利斯克VelskBarony struct {
	feud.BaseBarony
}

var BVelsk韦利斯克 feud.Barony = &韦利斯克VelskBarony{}

func init() {
    f := BVelsk韦利斯克.(*韦利斯克VelskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velsk",
		TitleName: "韦利斯克",
		TitleCode: "b_velsk",
	}
}
