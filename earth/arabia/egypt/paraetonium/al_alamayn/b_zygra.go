package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卒格拉ZygraBarony struct {
	feud.BaseBarony
}

var BZygra卒格拉 feud.Barony = &卒格拉ZygraBarony{}

func init() {
    f := BZygra卒格拉.(*卒格拉ZygraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zygra",
		TitleName: "卒格拉",
		TitleCode: "b_zygra",
	}
}
