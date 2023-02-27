package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达戈卢DagoluBarony struct {
	feud.BaseBarony
}

var BDagolu达戈卢 feud.Barony = &达戈卢DagoluBarony{}

func init() {
    f := BDagolu达戈卢.(*达戈卢DagoluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dagolu",
		TitleName: "达戈卢",
		TitleCode: "b_dagolu",
	}
}
