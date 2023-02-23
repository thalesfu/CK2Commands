package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔霍夫耶VerkhovyeBarony struct {
	feud.BaseBarony
}

var BVerkhovye韦尔霍夫耶 feud.Barony = &韦尔霍夫耶VerkhovyeBarony{}

func init() {
	f := BVerkhovye韦尔霍夫耶.(*韦尔霍夫耶VerkhovyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verkhovye",
		TitleName: "韦尔霍夫耶",
		TitleCode: "b_verkhovye",
	}
}
