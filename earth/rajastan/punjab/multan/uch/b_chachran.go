package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡克兰ChachranBarony struct {
	feud.BaseBarony
}

var BChachran卡克兰 feud.Barony = &卡克兰ChachranBarony{}

func init() {
    f := BChachran卡克兰.(*卡克兰ChachranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chachran",
		TitleName: "卡克兰",
		TitleCode: "b_chachran",
	}
}
