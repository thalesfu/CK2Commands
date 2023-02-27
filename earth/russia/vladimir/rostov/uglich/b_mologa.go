package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫洛加MologaBarony struct {
	feud.BaseBarony
}

var BMologa莫洛加 feud.Barony = &莫洛加MologaBarony{}

func init() {
    f := BMologa莫洛加.(*莫洛加MologaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mologa",
		TitleName: "莫洛加",
		TitleCode: "b_mologa",
	}
}
