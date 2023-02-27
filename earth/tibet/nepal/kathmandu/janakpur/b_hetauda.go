package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑道达HetaudaBarony struct {
	feud.BaseBarony
}

var BHetauda黑道达 feud.Barony = &黑道达HetaudaBarony{}

func init() {
    f := BHetauda黑道达.(*黑道达HetaudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hetauda",
		TitleName: "黑道达",
		TitleCode: "b_hetauda",
	}
}
