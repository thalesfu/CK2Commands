package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱卢厨摩隶KurudumaleBarony struct {
	feud.BaseBarony
}

var BKurudumale俱卢厨摩隶 feud.Barony = &俱卢厨摩隶KurudumaleBarony{}

func init() {
    f := BKurudumale俱卢厨摩隶.(*俱卢厨摩隶KurudumaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurudumale",
		TitleName: "俱卢厨摩隶",
		TitleCode: "b_kurudumale",
	}
}
