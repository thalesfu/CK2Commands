package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦里亚VeriaBarony struct {
	feud.BaseBarony
}

var BVeria韦里亚 feud.Barony = &韦里亚VeriaBarony{}

func init() {
	f := BVeria韦里亚.(*韦里亚VeriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veria",
		TitleName: "韦里亚",
		TitleCode: "b_veria",
	}
}
