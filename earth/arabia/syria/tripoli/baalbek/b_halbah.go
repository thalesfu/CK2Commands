package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈勒巴HalbahBarony struct {
	feud.BaseBarony
}

var BHalbah哈勒巴 feud.Barony = &哈勒巴HalbahBarony{}

func init() {
    f := BHalbah哈勒巴.(*哈勒巴HalbahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halbah",
		TitleName: "哈勒巴",
		TitleCode: "b_halbah",
	}
}
