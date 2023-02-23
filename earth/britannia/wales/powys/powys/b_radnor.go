package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德诺RadnorBarony struct {
	feud.BaseBarony
}

var BRadnor拉德诺 feud.Barony = &拉德诺RadnorBarony{}

func init() {
	f := BRadnor拉德诺.(*拉德诺RadnorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radnor",
		TitleName: "拉德诺",
		TitleCode: "b_radnor",
	}
}
