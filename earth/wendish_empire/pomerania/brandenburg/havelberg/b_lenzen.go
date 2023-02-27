package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦岑LenzenBarony struct {
	feud.BaseBarony
}

var BLenzen伦岑 feud.Barony = &伦岑LenzenBarony{}

func init() {
    f := BLenzen伦岑.(*伦岑LenzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lenzen",
		TitleName: "伦岑",
		TitleCode: "b_lenzen",
	}
}
