package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般滞湿伐罗PandeshwarBarony struct {
	feud.BaseBarony
}

var BPandeshwar般滞湿伐罗 feud.Barony = &般滞湿伐罗PandeshwarBarony{}

func init() {
	f := BPandeshwar般滞湿伐罗.(*般滞湿伐罗PandeshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pandeshwar",
		TitleName: "般滞湿伐罗",
		TitleCode: "b_pandeshwar",
	}
}
