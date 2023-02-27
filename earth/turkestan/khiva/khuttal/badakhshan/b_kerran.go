package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯朗KerranBarony struct {
	feud.BaseBarony
}

var BKerran凯朗 feud.Barony = &凯朗KerranBarony{}

func init() {
    f := BKerran凯朗.(*凯朗KerranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerran",
		TitleName: "凯朗",
		TitleCode: "b_kerran",
	}
}
