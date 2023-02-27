package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗伽吒BalaghatBarony struct {
	feud.BaseBarony
}

var BBalaghat婆罗伽吒 feud.Barony = &婆罗伽吒BalaghatBarony{}

func init() {
    f := BBalaghat婆罗伽吒.(*婆罗伽吒BalaghatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balaghat",
		TitleName: "婆罗伽吒",
		TitleCode: "b_balaghat",
	}
}
