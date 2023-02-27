package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗吉娑罗伊LakhisaraiBarony struct {
	feud.BaseBarony
}

var BLakhisarai罗吉娑罗伊 feud.Barony = &罗吉娑罗伊LakhisaraiBarony{}

func init() {
    f := BLakhisarai罗吉娑罗伊.(*罗吉娑罗伊LakhisaraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakhisarai",
		TitleName: "罗吉娑罗伊",
		TitleCode: "b_lakhisarai",
	}
}
