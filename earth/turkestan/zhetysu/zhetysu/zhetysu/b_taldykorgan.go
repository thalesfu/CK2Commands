package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔迪库尔干TaldykorganBarony struct {
	feud.BaseBarony
}

var BTaldykorgan塔尔迪库尔干 feud.Barony = &塔尔迪库尔干TaldykorganBarony{}

func init() {
	f := BTaldykorgan塔尔迪库尔干.(*塔尔迪库尔干TaldykorganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taldykorgan",
		TitleName: "塔尔迪库尔干",
		TitleCode: "b_taldykorgan",
	}
}
