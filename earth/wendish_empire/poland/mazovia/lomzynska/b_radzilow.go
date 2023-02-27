package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉济武夫RadzilowBarony struct {
	feud.BaseBarony
}

var BRadzilow拉济武夫 feud.Barony = &拉济武夫RadzilowBarony{}

func init() {
    f := BRadzilow拉济武夫.(*拉济武夫RadzilowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radzilow",
		TitleName: "拉济武夫",
		TitleCode: "b_radzilow",
	}
}
