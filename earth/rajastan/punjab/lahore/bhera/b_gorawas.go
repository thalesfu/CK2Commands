package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿罗婆斯GorawasBarony struct {
	feud.BaseBarony
}

var BGorawas瞿罗婆斯 feud.Barony = &瞿罗婆斯GorawasBarony{}

func init() {
	f := BGorawas瞿罗婆斯.(*瞿罗婆斯GorawasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorawas",
		TitleName: "瞿罗婆斯",
		TitleCode: "b_gorawas",
	}
}
