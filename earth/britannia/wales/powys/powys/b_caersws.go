package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔苏斯CaerswsBarony struct {
	feud.BaseBarony
}

var BCaersws凯尔苏斯 feud.Barony = &凯尔苏斯CaerswsBarony{}

func init() {
    f := BCaersws凯尔苏斯.(*凯尔苏斯CaerswsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caersws",
		TitleName: "凯尔苏斯",
		TitleCode: "b_caersws",
	}
}
