package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨雷斯CresBarony struct {
	feud.BaseBarony
}

var BCres茨雷斯 feud.Barony = &茨雷斯CresBarony{}

func init() {
    f := BCres茨雷斯.(*茨雷斯CresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cres",
		TitleName: "茨雷斯",
		TitleCode: "b_cres",
	}
}
