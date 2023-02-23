package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马希亚AmasiaBarony struct {
	feud.BaseBarony
}

var BAmasia阿马希亚 feud.Barony = &阿马希亚AmasiaBarony{}

func init() {
	f := BAmasia阿马希亚.(*阿马希亚AmasiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amasia",
		TitleName: "阿马希亚",
		TitleCode: "b_amasia",
	}
}
