package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔兰GulranBarony struct {
	feud.BaseBarony
}

var BGulran古尔兰 feud.Barony = &古尔兰GulranBarony{}

func init() {
    f := BGulran古尔兰.(*古尔兰GulranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gulran",
		TitleName: "古尔兰",
		TitleCode: "b_gulran",
	}
}
