package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡万德KavandBarony struct {
	feud.BaseBarony
}

var BKavand卡万德 feud.Barony = &卡万德KavandBarony{}

func init() {
    f := BKavand卡万德.(*卡万德KavandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kavand",
		TitleName: "卡万德",
		TitleCode: "b_kavand",
	}
}
