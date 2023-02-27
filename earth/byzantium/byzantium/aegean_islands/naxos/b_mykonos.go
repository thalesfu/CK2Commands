package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米科诺斯MykonosBarony struct {
	feud.BaseBarony
}

var BMykonos米科诺斯 feud.Barony = &米科诺斯MykonosBarony{}

func init() {
    f := BMykonos米科诺斯.(*米科诺斯MykonosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mykonos",
		TitleName: "米科诺斯",
		TitleCode: "b_mykonos",
	}
}
