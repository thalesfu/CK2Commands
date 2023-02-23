package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法恩扎FaenzaBarony struct {
	feud.BaseBarony
}

var BFaenza法恩扎 feud.Barony = &法恩扎FaenzaBarony{}

func init() {
	f := BFaenza法恩扎.(*法恩扎FaenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faenza",
		TitleName: "法恩扎",
		TitleCode: "b_faenza",
	}
}
