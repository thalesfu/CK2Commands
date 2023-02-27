package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟索ThursoBarony struct {
	feud.BaseBarony
}

var BThurso瑟索 feud.Barony = &瑟索ThursoBarony{}

func init() {
    f := BThurso瑟索.(*瑟索ThursoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thurso",
		TitleName: "瑟索",
		TitleCode: "b_thurso",
	}
}
