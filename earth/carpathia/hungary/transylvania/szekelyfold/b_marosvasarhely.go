package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛罗什瓦沙尔海伊MarosvasarhelyBarony struct {
	feud.BaseBarony
}

var BMarosvasarhely毛罗什瓦沙尔海伊 feud.Barony = &毛罗什瓦沙尔海伊MarosvasarhelyBarony{}

func init() {
    f := BMarosvasarhely毛罗什瓦沙尔海伊.(*毛罗什瓦沙尔海伊MarosvasarhelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marosvasarhely",
		TitleName: "毛罗什瓦沙尔海伊",
		TitleCode: "b_marosvasarhely",
	}
}
