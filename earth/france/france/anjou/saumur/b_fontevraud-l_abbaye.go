package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰特夫罗拉拜Fontevraud_l_abbayeBarony struct {
	feud.BaseBarony
}

var BFontevraud_l_abbaye丰特夫罗拉拜 feud.Barony = &丰特夫罗拉拜Fontevraud_l_abbayeBarony{}

func init() {
    f := BFontevraud_l_abbaye丰特夫罗拉拜.(*丰特夫罗拉拜Fontevraud_l_abbayeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fontevraud_l_abbaye",
		TitleName: "丰特夫罗拉拜",
		TitleCode: "b_fontevraud-l_abbaye",
	}
}
