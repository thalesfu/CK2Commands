package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里格BrigBarony struct {
	feud.BaseBarony
}

var BBrig布里格 feud.Barony = &布里格BrigBarony{}

func init() {
    f := BBrig布里格.(*布里格BrigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brig",
		TitleName: "布里格",
		TitleCode: "b_brig",
	}
}
