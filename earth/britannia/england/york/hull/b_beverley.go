package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝弗利BeverleyBarony struct {
	feud.BaseBarony
}

var BBeverley贝弗利 feud.Barony = &贝弗利BeverleyBarony{}

func init() {
	f := BBeverley贝弗利.(*贝弗利BeverleyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beverley",
		TitleName: "贝弗利",
		TitleCode: "b_beverley",
	}
}
