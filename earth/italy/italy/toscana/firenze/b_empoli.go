package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩波利EmpoliBarony struct {
	feud.BaseBarony
}

var BEmpoli恩波利 feud.Barony = &恩波利EmpoliBarony{}

func init() {
    f := BEmpoli恩波利.(*恩波利EmpoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "empoli",
		TitleName: "恩波利",
		TitleCode: "b_empoli",
	}
}
