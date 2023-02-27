package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钵吒那拘吒PathankotBarony struct {
	feud.BaseBarony
}

var BPathankot钵吒那拘吒 feud.Barony = &钵吒那拘吒PathankotBarony{}

func init() {
    f := BPathankot钵吒那拘吒.(*钵吒那拘吒PathankotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pathankot",
		TitleName: "钵吒那拘吒",
		TitleCode: "b_pathankot",
	}
}
