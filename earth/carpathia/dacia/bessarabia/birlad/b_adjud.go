package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德茹德AdjudBarony struct {
	feud.BaseBarony
}

var BAdjud阿德茹德 feud.Barony = &阿德茹德AdjudBarony{}

func init() {
	f := BAdjud阿德茹德.(*阿德茹德AdjudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adjud",
		TitleName: "阿德茹德",
		TitleCode: "b_adjud",
	}
}
