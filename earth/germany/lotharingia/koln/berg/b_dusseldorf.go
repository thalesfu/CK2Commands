package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜塞尔多夫DusseldorfBarony struct {
	feud.BaseBarony
}

var BDusseldorf杜塞尔多夫 feud.Barony = &杜塞尔多夫DusseldorfBarony{}

func init() {
    f := BDusseldorf杜塞尔多夫.(*杜塞尔多夫DusseldorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dusseldorf",
		TitleName: "杜塞尔多夫",
		TitleCode: "b_dusseldorf",
	}
}
