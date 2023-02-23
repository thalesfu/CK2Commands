package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布克BukBarony struct {
	feud.BaseBarony
}

var BBuk布克 feud.Barony = &布克BukBarony{}

func init() {
	f := BBuk布克.(*布克BukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buk",
		TitleName: "布克",
		TitleCode: "b_buk",
	}
}
