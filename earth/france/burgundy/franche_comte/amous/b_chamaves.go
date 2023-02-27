package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨马维斯ChamavesBarony struct {
	feud.BaseBarony
}

var BChamaves萨马维斯 feud.Barony = &萨马维斯ChamavesBarony{}

func init() {
    f := BChamaves萨马维斯.(*萨马维斯ChamavesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chamaves",
		TitleName: "萨马维斯",
		TitleCode: "b_chamaves",
	}
}
