package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯特莫EstemonBarony struct {
	feud.BaseBarony
}

var BEstemon埃斯特莫 feud.Barony = &埃斯特莫EstemonBarony{}

func init() {
    f := BEstemon埃斯特莫.(*埃斯特莫EstemonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estemon",
		TitleName: "埃斯特莫",
		TitleCode: "b_estemon",
	}
}
