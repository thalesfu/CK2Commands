package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯图泽尼采StudzienikiBarony struct {
	feud.BaseBarony
}

var BStudzieniki斯图泽尼采 feud.Barony = &斯图泽尼采StudzienikiBarony{}

func init() {
    f := BStudzieniki斯图泽尼采.(*斯图泽尼采StudzienikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "studzieniki",
		TitleName: "斯图泽尼采",
		TitleCode: "b_studzieniki",
	}
}
