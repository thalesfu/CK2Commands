package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔布克TabukBarony struct {
	feud.BaseBarony
}

var BTabuk塔布克 feud.Barony = &塔布克TabukBarony{}

func init() {
    f := BTabuk塔布克.(*塔布克TabukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabuk",
		TitleName: "塔布克",
		TitleCode: "b_tabuk",
	}
}
