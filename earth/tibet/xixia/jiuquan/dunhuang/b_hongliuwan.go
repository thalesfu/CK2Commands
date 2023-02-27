package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 红柳湾HongliuwanBarony struct {
	feud.BaseBarony
}

var BHongliuwan红柳湾 feud.Barony = &红柳湾HongliuwanBarony{}

func init() {
    f := BHongliuwan红柳湾.(*红柳湾HongliuwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hongliuwan",
		TitleName: "红柳湾",
		TitleCode: "b_hongliuwan",
	}
}
