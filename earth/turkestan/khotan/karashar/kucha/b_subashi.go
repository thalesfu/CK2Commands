package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏巴什SubashiBarony struct {
	feud.BaseBarony
}

var BSubashi苏巴什 feud.Barony = &苏巴什SubashiBarony{}

func init() {
    f := BSubashi苏巴什.(*苏巴什SubashiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "subashi",
		TitleName: "苏巴什",
		TitleCode: "b_subashi",
	}
}
