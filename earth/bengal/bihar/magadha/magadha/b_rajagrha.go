package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 王舍城RajagrhaBarony struct {
	feud.BaseBarony
}

var BRajagrha王舍城 feud.Barony = &王舍城RajagrhaBarony{}

func init() {
    f := BRajagrha王舍城.(*王舍城RajagrhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajagrha",
		TitleName: "王舍城",
		TitleCode: "b_rajagrha",
	}
}
