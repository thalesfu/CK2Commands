package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗霍伊DorohoiBarony struct {
	feud.BaseBarony
}

var BDorohoi多罗霍伊 feud.Barony = &多罗霍伊DorohoiBarony{}

func init() {
    f := BDorohoi多罗霍伊.(*多罗霍伊DorohoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorohoi",
		TitleName: "多罗霍伊",
		TitleCode: "b_dorohoi",
	}
}
