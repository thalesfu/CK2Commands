package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼奴怛婆提NibnutvadiBarony struct {
	feud.BaseBarony
}

var BNibnutvadi尼奴怛婆提 feud.Barony = &尼奴怛婆提NibnutvadiBarony{}

func init() {
    f := BNibnutvadi尼奴怛婆提.(*尼奴怛婆提NibnutvadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nibnutvadi",
		TitleName: "尼奴怛婆提",
		TitleCode: "b_nibnutvadi",
	}
}
