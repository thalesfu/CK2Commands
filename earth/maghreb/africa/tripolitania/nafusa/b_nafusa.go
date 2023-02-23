package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈富塞NafusaBarony struct {
	feud.BaseBarony
}

var BNafusa奈富塞 feud.Barony = &奈富塞NafusaBarony{}

func init() {
	f := BNafusa奈富塞.(*奈富塞NafusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nafusa",
		TitleName: "奈富塞",
		TitleCode: "b_nafusa",
	}
}
