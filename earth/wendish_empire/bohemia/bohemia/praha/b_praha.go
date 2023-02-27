package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉格PrahaBarony struct {
	feud.BaseBarony
}

var BPraha布拉格 feud.Barony = &布拉格PrahaBarony{}

func init() {
    f := BPraha布拉格.(*布拉格PrahaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "praha",
		TitleName: "布拉格",
		TitleCode: "b_praha",
	}
}
