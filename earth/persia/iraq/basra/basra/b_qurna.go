package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀尔那QurnaBarony struct {
	feud.BaseBarony
}

var BQurna喀尔那 feud.Barony = &喀尔那QurnaBarony{}

func init() {
    f := BQurna喀尔那.(*喀尔那QurnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qurna",
		TitleName: "喀尔那",
		TitleCode: "b_qurna",
	}
}
