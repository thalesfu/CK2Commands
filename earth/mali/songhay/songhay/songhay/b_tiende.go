package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂安德TiendeBarony struct {
	feud.BaseBarony
}

var BTiende蒂安德 feud.Barony = &蒂安德TiendeBarony{}

func init() {
	f := BTiende蒂安德.(*蒂安德TiendeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiende",
		TitleName: "蒂安德",
		TitleCode: "b_tiende",
	}
}
