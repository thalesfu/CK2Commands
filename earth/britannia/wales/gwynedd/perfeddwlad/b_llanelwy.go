package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰埃尔威LlanelwyBarony struct {
	feud.BaseBarony
}

var BLlanelwy兰埃尔威 feud.Barony = &兰埃尔威LlanelwyBarony{}

func init() {
    f := BLlanelwy兰埃尔威.(*兰埃尔威LlanelwyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanelwy",
		TitleName: "兰埃尔威",
		TitleCode: "b_llanelwy",
	}
}
