package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼姆NimesBarony struct {
	feud.BaseBarony
}

var BNimes尼姆 feud.Barony = &尼姆NimesBarony{}

func init() {
    f := BNimes尼姆.(*尼姆NimesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimes",
		TitleName: "尼姆",
		TitleCode: "b_nimes",
	}
}
