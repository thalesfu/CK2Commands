package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩普里耶斯EmpuriesBarony struct {
	feud.BaseBarony
}

var BEmpuries恩普里耶斯 feud.Barony = &恩普里耶斯EmpuriesBarony{}

func init() {
    f := BEmpuries恩普里耶斯.(*恩普里耶斯EmpuriesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "empuries",
		TitleName: "恩普里耶斯",
		TitleCode: "b_empuries",
	}
}
