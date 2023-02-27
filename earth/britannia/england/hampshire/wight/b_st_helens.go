package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣海伦斯St_helensBarony struct {
	feud.BaseBarony
}

var BSt_helens圣海伦斯 feud.Barony = &圣海伦斯St_helensBarony{}

func init() {
    f := BSt_helens圣海伦斯.(*圣海伦斯St_helensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_helens",
		TitleName: "圣海伦斯",
		TitleCode: "b_st_helens",
	}
}
