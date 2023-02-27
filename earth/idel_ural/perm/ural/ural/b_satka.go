package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨特卡SatkaBarony struct {
	feud.BaseBarony
}

var BSatka萨特卡 feud.Barony = &萨特卡SatkaBarony{}

func init() {
    f := BSatka萨特卡.(*萨特卡SatkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satka",
		TitleName: "萨特卡",
		TitleCode: "b_satka",
	}
}
