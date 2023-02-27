package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代特莫尔德DetmoldBarony struct {
	feud.BaseBarony
}

var BDetmold代特莫尔德 feud.Barony = &代特莫尔德DetmoldBarony{}

func init() {
    f := BDetmold代特莫尔德.(*代特莫尔德DetmoldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "detmold",
		TitleName: "代特莫尔德",
		TitleCode: "b_detmold",
	}
}
