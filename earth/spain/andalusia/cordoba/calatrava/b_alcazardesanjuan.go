package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣胡安堡AlcazardesanjuanBarony struct {
	feud.BaseBarony
}

var BAlcazardesanjuan圣胡安堡 feud.Barony = &圣胡安堡AlcazardesanjuanBarony{}

func init() {
	f := BAlcazardesanjuan圣胡安堡.(*圣胡安堡AlcazardesanjuanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcazardesanjuan",
		TitleName: "圣胡安堡",
		TitleCode: "b_alcazardesanjuan",
	}
}
