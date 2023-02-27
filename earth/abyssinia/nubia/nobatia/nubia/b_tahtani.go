package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫塔尼TahtaniBarony struct {
	feud.BaseBarony
}

var BTahtani塔赫塔尼 feud.Barony = &塔赫塔尼TahtaniBarony{}

func init() {
    f := BTahtani塔赫塔尼.(*塔赫塔尼TahtaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tahtani",
		TitleName: "塔赫塔尼",
		TitleCode: "b_tahtani",
	}
}
