package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钵陀劳那PadraunaBarony struct {
	feud.BaseBarony
}

var BPadrauna钵陀劳那 feud.Barony = &钵陀劳那PadraunaBarony{}

func init() {
    f := BPadrauna钵陀劳那.(*钵陀劳那PadraunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "padrauna",
		TitleName: "钵陀劳那",
		TitleCode: "b_padrauna",
	}
}
