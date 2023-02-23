package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费里亚PhaliaBarony struct {
	feud.BaseBarony
}

var BPhalia费里亚 feud.Barony = &费里亚PhaliaBarony{}

func init() {
	f := BPhalia费里亚.(*费里亚PhaliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phalia",
		TitleName: "费里亚",
		TitleCode: "b_phalia",
	}
}
