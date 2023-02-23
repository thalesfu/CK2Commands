package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格达BagdadBarony struct {
	feud.BaseBarony
}

var BBagdad巴格达 feud.Barony = &巴格达BagdadBarony{}

func init() {
	f := BBagdad巴格达.(*巴格达BagdadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagdad",
		TitleName: "巴格达",
		TitleCode: "b_bagdad",
	}
}
