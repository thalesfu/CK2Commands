package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔热祖尔AljezurBarony struct {
	feud.BaseBarony
}

var BAljezur阿尔热祖尔 feud.Barony = &阿尔热祖尔AljezurBarony{}

func init() {
    f := BAljezur阿尔热祖尔.(*阿尔热祖尔AljezurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljezur",
		TitleName: "阿尔热祖尔",
		TitleCode: "b_aljezur",
	}
}
