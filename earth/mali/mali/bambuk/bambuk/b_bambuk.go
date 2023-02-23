package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班布克BambukBarony struct {
	feud.BaseBarony
}

var BBambuk班布克 feud.Barony = &班布克BambukBarony{}

func init() {
	f := BBambuk班布克.(*班布克BambukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bambuk",
		TitleName: "班布克",
		TitleCode: "b_bambuk",
	}
}
