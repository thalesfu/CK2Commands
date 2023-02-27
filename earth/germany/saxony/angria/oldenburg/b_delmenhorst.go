package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代尔门霍斯特DelmenhorstBarony struct {
	feud.BaseBarony
}

var BDelmenhorst代尔门霍斯特 feud.Barony = &代尔门霍斯特DelmenhorstBarony{}

func init() {
    f := BDelmenhorst代尔门霍斯特.(*代尔门霍斯特DelmenhorstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delmenhorst",
		TitleName: "代尔门霍斯特",
		TitleCode: "b_delmenhorst",
	}
}
