package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅乌赤YauchyBarony struct {
	feud.BaseBarony
}

var BYauchy雅乌赤 feud.Barony = &雅乌赤YauchyBarony{}

func init() {
    f := BYauchy雅乌赤.(*雅乌赤YauchyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yauchy",
		TitleName: "雅乌赤",
		TitleCode: "b_yauchy",
	}
}
