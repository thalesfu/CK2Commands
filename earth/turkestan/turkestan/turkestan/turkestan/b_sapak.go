package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨帕克SapakBarony struct {
	feud.BaseBarony
}

var BSapak萨帕克 feud.Barony = &萨帕克SapakBarony{}

func init() {
    f := BSapak萨帕克.(*萨帕克SapakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sapak",
		TitleName: "萨帕克",
		TitleCode: "b_sapak",
	}
}
