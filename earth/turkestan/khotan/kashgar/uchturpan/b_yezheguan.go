package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谒者馆YezheguanBarony struct {
	feud.BaseBarony
}

var BYezheguan谒者馆 feud.Barony = &谒者馆YezheguanBarony{}

func init() {
    f := BYezheguan谒者馆.(*谒者馆YezheguanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yezheguan",
		TitleName: "谒者馆",
		TitleCode: "b_yezheguan",
	}
}
