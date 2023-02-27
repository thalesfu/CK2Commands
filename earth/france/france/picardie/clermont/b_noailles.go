package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺阿耶NoaillesBarony struct {
	feud.BaseBarony
}

var BNoailles诺阿耶 feud.Barony = &诺阿耶NoaillesBarony{}

func init() {
    f := BNoailles诺阿耶.(*诺阿耶NoaillesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noailles",
		TitleName: "诺阿耶",
		TitleCode: "b_noailles",
	}
}
