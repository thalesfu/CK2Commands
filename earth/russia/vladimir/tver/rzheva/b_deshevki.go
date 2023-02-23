package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷舍夫基DeshevkiBarony struct {
	feud.BaseBarony
}

var BDeshevki捷舍夫基 feud.Barony = &捷舍夫基DeshevkiBarony{}

func init() {
	f := BDeshevki捷舍夫基.(*捷舍夫基DeshevkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deshevki",
		TitleName: "捷舍夫基",
		TitleCode: "b_deshevki",
	}
}
