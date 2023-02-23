package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔迪卡SerdicaBarony struct {
	feud.BaseBarony
}

var BSerdica塞尔迪卡 feud.Barony = &塞尔迪卡SerdicaBarony{}

func init() {
	f := BSerdica塞尔迪卡.(*塞尔迪卡SerdicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serdica",
		TitleName: "塞尔迪卡",
		TitleCode: "b_serdica",
	}
}
