package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔顿BoltonBarony struct {
	feud.BaseBarony
}

var BBolton博尔顿 feud.Barony = &博尔顿BoltonBarony{}

func init() {
    f := BBolton博尔顿.(*博尔顿BoltonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolton",
		TitleName: "博尔顿",
		TitleCode: "b_bolton",
	}
}
