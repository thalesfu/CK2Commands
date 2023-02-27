package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格勒诺布尔GrenobleBarony struct {
	feud.BaseBarony
}

var BGrenoble格勒诺布尔 feud.Barony = &格勒诺布尔GrenobleBarony{}

func init() {
    f := BGrenoble格勒诺布尔.(*格勒诺布尔GrenobleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grenoble",
		TitleName: "格勒诺布尔",
		TitleCode: "b_grenoble",
	}
}
