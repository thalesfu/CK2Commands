package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾勒格JaleqBarony struct {
	feud.BaseBarony
}

var BJaleq贾勒格 feud.Barony = &贾勒格JaleqBarony{}

func init() {
	f := BJaleq贾勒格.(*贾勒格JaleqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaleq",
		TitleName: "贾勒格",
		TitleCode: "b_jaleq",
	}
}
