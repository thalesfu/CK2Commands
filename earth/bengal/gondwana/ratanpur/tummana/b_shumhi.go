package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 首摩醯ShumhiBarony struct {
	feud.BaseBarony
}

var BShumhi首摩醯 feud.Barony = &首摩醯ShumhiBarony{}

func init() {
    f := BShumhi首摩醯.(*首摩醯ShumhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shumhi",
		TitleName: "首摩醯",
		TitleCode: "b_shumhi",
	}
}
