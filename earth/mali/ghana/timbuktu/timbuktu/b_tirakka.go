package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂拉卡TirakkaBarony struct {
	feud.BaseBarony
}

var BTirakka蒂拉卡 feud.Barony = &蒂拉卡TirakkaBarony{}

func init() {
    f := BTirakka蒂拉卡.(*蒂拉卡TirakkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirakka",
		TitleName: "蒂拉卡",
		TitleCode: "b_tirakka",
	}
}
