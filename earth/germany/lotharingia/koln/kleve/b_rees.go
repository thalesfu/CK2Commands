package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷斯ReesBarony struct {
	feud.BaseBarony
}

var BRees雷斯 feud.Barony = &雷斯ReesBarony{}

func init() {
    f := BRees雷斯.(*雷斯ReesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rees",
		TitleName: "雷斯",
		TitleCode: "b_rees",
	}
}
