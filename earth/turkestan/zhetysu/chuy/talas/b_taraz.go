package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 怛逻斯TarazBarony struct {
	feud.BaseBarony
}

var BTaraz怛逻斯 feud.Barony = &怛逻斯TarazBarony{}

func init() {
    f := BTaraz怛逻斯.(*怛逻斯TarazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taraz",
		TitleName: "怛逻斯",
		TitleCode: "b_taraz",
	}
}
