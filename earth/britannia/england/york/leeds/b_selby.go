package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔比SelbyBarony struct {
	feud.BaseBarony
}

var BSelby塞尔比 feud.Barony = &塞尔比SelbyBarony{}

func init() {
    f := BSelby塞尔比.(*塞尔比SelbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selby",
		TitleName: "塞尔比",
		TitleCode: "b_selby",
	}
}
