package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格兰纳GrennahBarony struct {
	feud.BaseBarony
}

var BGrennah格兰纳 feud.Barony = &格兰纳GrennahBarony{}

func init() {
    f := BGrennah格兰纳.(*格兰纳GrennahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grennah",
		TitleName: "格兰纳",
		TitleCode: "b_grennah",
	}
}
