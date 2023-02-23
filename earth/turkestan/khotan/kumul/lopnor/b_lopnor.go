package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗布泊LopnorBarony struct {
	feud.BaseBarony
}

var BLopnor罗布泊 feud.Barony = &罗布泊LopnorBarony{}

func init() {
	f := BLopnor罗布泊.(*罗布泊LopnorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lopnor",
		TitleName: "罗布泊",
		TitleCode: "b_lopnor",
	}
}
