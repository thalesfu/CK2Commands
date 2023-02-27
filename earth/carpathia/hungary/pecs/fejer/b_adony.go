package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥多尼AdonyBarony struct {
	feud.BaseBarony
}

var BAdony奥多尼 feud.Barony = &奥多尼AdonyBarony{}

func init() {
    f := BAdony奥多尼.(*奥多尼AdonyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adony",
		TitleName: "奥多尼",
		TitleCode: "b_adony",
	}
}
