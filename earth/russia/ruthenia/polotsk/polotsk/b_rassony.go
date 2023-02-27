package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗索内RassonyBarony struct {
	feud.BaseBarony
}

var BRassony罗索内 feud.Barony = &罗索内RassonyBarony{}

func init() {
    f := BRassony罗索内.(*罗索内RassonyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rassony",
		TitleName: "罗索内",
		TitleCode: "b_rassony",
	}
}
