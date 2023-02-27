package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔赫科塔尔Surkh_kotalBarony struct {
	feud.BaseBarony
}

var BSurkh_kotal苏尔赫科塔尔 feud.Barony = &苏尔赫科塔尔Surkh_kotalBarony{}

func init() {
    f := BSurkh_kotal苏尔赫科塔尔.(*苏尔赫科塔尔Surkh_kotalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surkh_kotal",
		TitleName: "苏尔赫科塔尔",
		TitleCode: "b_surkh_kotal",
	}
}
