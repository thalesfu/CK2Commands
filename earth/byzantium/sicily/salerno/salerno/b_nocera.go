package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺切拉NoceraBarony struct {
	feud.BaseBarony
}

var BNocera诺切拉 feud.Barony = &诺切拉NoceraBarony{}

func init() {
    f := BNocera诺切拉.(*诺切拉NoceraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nocera",
		TitleName: "诺切拉",
		TitleCode: "b_nocera",
	}
}
