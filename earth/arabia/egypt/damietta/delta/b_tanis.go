package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尼斯TanisBarony struct {
	feud.BaseBarony
}

var BTanis塔尼斯 feud.Barony = &塔尼斯TanisBarony{}

func init() {
    f := BTanis塔尼斯.(*塔尼斯TanisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanis",
		TitleName: "塔尼斯",
		TitleCode: "b_tanis",
	}
}
