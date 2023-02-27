package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃克塞特ExeterBarony struct {
	feud.BaseBarony
}

var BExeter埃克塞特 feud.Barony = &埃克塞特ExeterBarony{}

func init() {
    f := BExeter埃克塞特.(*埃克塞特ExeterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "exeter",
		TitleName: "埃克塞特",
		TitleCode: "b_exeter",
	}
}
