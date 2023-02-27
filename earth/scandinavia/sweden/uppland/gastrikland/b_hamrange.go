package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈姆罗耶HamrangeBarony struct {
	feud.BaseBarony
}

var BHamrange哈姆罗耶 feud.Barony = &哈姆罗耶HamrangeBarony{}

func init() {
    f := BHamrange哈姆罗耶.(*哈姆罗耶HamrangeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamrange",
		TitleName: "哈姆罗耶",
		TitleCode: "b_hamrange",
	}
}
