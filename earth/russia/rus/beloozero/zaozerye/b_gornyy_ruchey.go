package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔内鲁切伊Gornyy_rucheyBarony struct {
	feud.BaseBarony
}

var BGornyy_ruchey戈尔内鲁切伊 feud.Barony = &戈尔内鲁切伊Gornyy_rucheyBarony{}

func init() {
    f := BGornyy_ruchey戈尔内鲁切伊.(*戈尔内鲁切伊Gornyy_rucheyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gornyy_ruchey",
		TitleName: "戈尔内鲁切伊",
		TitleCode: "b_gornyy_ruchey",
	}
}
