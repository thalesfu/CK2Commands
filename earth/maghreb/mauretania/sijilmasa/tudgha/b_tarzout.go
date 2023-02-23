package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔祖特TarzoutBarony struct {
	feud.BaseBarony
}

var BTarzout塔尔祖特 feud.Barony = &塔尔祖特TarzoutBarony{}

func init() {
	f := BTarzout塔尔祖特.(*塔尔祖特TarzoutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarzout",
		TitleName: "塔尔祖特",
		TitleCode: "b_tarzout",
	}
}
