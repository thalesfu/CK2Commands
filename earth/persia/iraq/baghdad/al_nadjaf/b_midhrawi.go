package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米德霍拉维MidhrawiBarony struct {
	feud.BaseBarony
}

var BMidhrawi米德霍拉维 feud.Barony = &米德霍拉维MidhrawiBarony{}

func init() {
    f := BMidhrawi米德霍拉维.(*米德霍拉维MidhrawiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "midhrawi",
		TitleName: "米德霍拉维",
		TitleCode: "b_midhrawi",
	}
}
