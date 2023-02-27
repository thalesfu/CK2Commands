package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗马贾凯Horo_majakaiBarony struct {
	feud.BaseBarony
}

var BHoro_majakai霍罗马贾凯 feud.Barony = &霍罗马贾凯Horo_majakaiBarony{}

func init() {
    f := BHoro_majakai霍罗马贾凯.(*霍罗马贾凯Horo_majakaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horo_majakai",
		TitleName: "霍罗马贾凯",
		TitleCode: "b_horo_majakai",
	}
}
