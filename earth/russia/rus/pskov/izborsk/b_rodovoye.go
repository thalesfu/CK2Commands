package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗多沃耶RodovoyeBarony struct {
	feud.BaseBarony
}

var BRodovoye罗多沃耶 feud.Barony = &罗多沃耶RodovoyeBarony{}

func init() {
	f := BRodovoye罗多沃耶.(*罗多沃耶RodovoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodovoye",
		TitleName: "罗多沃耶",
		TitleCode: "b_rodovoye",
	}
}
