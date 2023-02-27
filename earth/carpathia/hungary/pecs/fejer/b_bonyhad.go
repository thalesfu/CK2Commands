package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尼哈德BonyhadBarony struct {
	feud.BaseBarony
}

var BBonyhad博尼哈德 feud.Barony = &博尼哈德BonyhadBarony{}

func init() {
    f := BBonyhad博尼哈德.(*博尼哈德BonyhadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonyhad",
		TitleName: "博尼哈德",
		TitleCode: "b_bonyhad",
	}
}
