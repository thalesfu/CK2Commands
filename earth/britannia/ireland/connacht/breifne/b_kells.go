package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔斯KellsBarony struct {
	feud.BaseBarony
}

var BKells凯尔斯 feud.Barony = &凯尔斯KellsBarony{}

func init() {
    f := BKells凯尔斯.(*凯尔斯KellsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kells",
		TitleName: "凯尔斯",
		TitleCode: "b_kells",
	}
}
