package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特布鲁赫OsterbruchBarony struct {
	feud.BaseBarony
}

var BOsterbruch奥斯特布鲁赫 feud.Barony = &奥斯特布鲁赫OsterbruchBarony{}

func init() {
	f := BOsterbruch奥斯特布鲁赫.(*奥斯特布鲁赫OsterbruchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osterbruch",
		TitleName: "奥斯特布鲁赫",
		TitleCode: "b_osterbruch",
	}
}
