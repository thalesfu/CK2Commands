package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔布兰KerburanBarony struct {
	feud.BaseBarony
}

var BKerburan克尔布兰 feud.Barony = &克尔布兰KerburanBarony{}

func init() {
	f := BKerburan克尔布兰.(*克尔布兰KerburanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerburan",
		TitleName: "克尔布兰",
		TitleCode: "b_kerburan",
	}
}
