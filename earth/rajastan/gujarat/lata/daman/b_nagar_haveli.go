package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那迦诃伐尼Nagar_haveliBarony struct {
	feud.BaseBarony
}

var BNagar_haveli那迦诃伐尼 feud.Barony = &那迦诃伐尼Nagar_haveliBarony{}

func init() {
    f := BNagar_haveli那迦诃伐尼.(*那迦诃伐尼Nagar_haveliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagar_haveli",
		TitleName: "那迦诃伐尼",
		TitleCode: "b_nagar_haveli",
	}
}
