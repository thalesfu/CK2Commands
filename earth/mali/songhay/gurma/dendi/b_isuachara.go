package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊苏阿查拉IsuacharaBarony struct {
	feud.BaseBarony
}

var BIsuachara伊苏阿查拉 feud.Barony = &伊苏阿查拉IsuacharaBarony{}

func init() {
    f := BIsuachara伊苏阿查拉.(*伊苏阿查拉IsuacharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isuachara",
		TitleName: "伊苏阿查拉",
		TitleCode: "b_isuachara",
	}
}
