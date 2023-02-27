package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉芬斯泰因RavensteinBarony struct {
	feud.BaseBarony
}

var BRavenstein拉芬斯泰因 feud.Barony = &拉芬斯泰因RavensteinBarony{}

func init() {
    f := BRavenstein拉芬斯泰因.(*拉芬斯泰因RavensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravenstein",
		TitleName: "拉芬斯泰因",
		TitleCode: "b_ravenstein",
	}
}
