package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦杰诺VedenoBarony struct {
	feud.BaseBarony
}

var BVedeno韦杰诺 feud.Barony = &韦杰诺VedenoBarony{}

func init() {
    f := BVedeno韦杰诺.(*韦杰诺VedenoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vedeno",
		TitleName: "韦杰诺",
		TitleCode: "b_vedeno",
	}
}
