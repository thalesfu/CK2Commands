package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费德克斯腾FredrikstenBarony struct {
	feud.BaseBarony
}

var BFredriksten费德克斯腾 feud.Barony = &费德克斯腾FredrikstenBarony{}

func init() {
    f := BFredriksten费德克斯腾.(*费德克斯腾FredrikstenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fredriksten",
		TitleName: "费德克斯腾",
		TitleCode: "b_fredriksten",
	}
}
