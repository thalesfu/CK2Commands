package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉古萨RagusaBarony struct {
	feud.BaseBarony
}

var BRagusa拉古萨 feud.Barony = &拉古萨RagusaBarony{}

func init() {
    f := BRagusa拉古萨.(*拉古萨RagusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ragusa",
		TitleName: "拉古萨",
		TitleCode: "b_ragusa",
	}
}
