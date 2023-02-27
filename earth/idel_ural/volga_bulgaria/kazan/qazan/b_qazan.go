package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀山QazanBarony struct {
	feud.BaseBarony
}

var BQazan喀山 feud.Barony = &喀山QazanBarony{}

func init() {
    f := BQazan喀山.(*喀山QazanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qazan",
		TitleName: "喀山",
		TitleCode: "b_qazan",
	}
}
