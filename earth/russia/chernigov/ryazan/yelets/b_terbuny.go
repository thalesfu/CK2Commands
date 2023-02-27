package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷尔布内TerbunyBarony struct {
	feud.BaseBarony
}

var BTerbuny捷尔布内 feud.Barony = &捷尔布内TerbunyBarony{}

func init() {
    f := BTerbuny捷尔布内.(*捷尔布内TerbunyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terbuny",
		TitleName: "捷尔布内",
		TitleCode: "b_terbuny",
	}
}
