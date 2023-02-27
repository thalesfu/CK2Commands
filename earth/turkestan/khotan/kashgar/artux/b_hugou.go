package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湖沟HugouBarony struct {
	feud.BaseBarony
}

var BHugou湖沟 feud.Barony = &湖沟HugouBarony{}

func init() {
    f := BHugou湖沟.(*湖沟HugouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hugou",
		TitleName: "湖沟",
		TitleCode: "b_hugou",
	}
}
