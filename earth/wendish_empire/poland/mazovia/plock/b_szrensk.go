package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什伦斯克SzrenskBarony struct {
	feud.BaseBarony
}

var BSzrensk什伦斯克 feud.Barony = &什伦斯克SzrenskBarony{}

func init() {
    f := BSzrensk什伦斯克.(*什伦斯克SzrenskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szrensk",
		TitleName: "什伦斯克",
		TitleCode: "b_szrensk",
	}
}
