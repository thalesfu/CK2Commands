package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博赫尼亚BochniaBarony struct {
	feud.BaseBarony
}

var BBochnia博赫尼亚 feud.Barony = &博赫尼亚BochniaBarony{}

func init() {
    f := BBochnia博赫尼亚.(*博赫尼亚BochniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bochnia",
		TitleName: "博赫尼亚",
		TitleCode: "b_bochnia",
	}
}
