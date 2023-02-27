package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里涅BrinjeBarony struct {
	feud.BaseBarony
}

var BBrinje布里涅 feud.Barony = &布里涅BrinjeBarony{}

func init() {
    f := BBrinje布里涅.(*布里涅BrinjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brinje",
		TitleName: "布里涅",
		TitleCode: "b_brinje",
	}
}
