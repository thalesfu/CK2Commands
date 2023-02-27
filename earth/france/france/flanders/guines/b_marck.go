package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马克MarckBarony struct {
	feud.BaseBarony
}

var BMarck马克 feud.Barony = &马克MarckBarony{}

func init() {
    f := BMarck马克.(*马克MarckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marck",
		TitleName: "马克",
		TitleCode: "b_marck",
	}
}
