package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特伦AthrunBarony struct {
	feud.BaseBarony
}

var BAthrun阿特伦 feud.Barony = &阿特伦AthrunBarony{}

func init() {
    f := BAthrun阿特伦.(*阿特伦AthrunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "athrun",
		TitleName: "阿特伦",
		TitleCode: "b_athrun",
	}
}
