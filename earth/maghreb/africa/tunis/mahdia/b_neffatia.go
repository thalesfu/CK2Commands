package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒法蒂亚NeffatiaBarony struct {
	feud.BaseBarony
}

var BNeffatia勒法蒂亚 feud.Barony = &勒法蒂亚NeffatiaBarony{}

func init() {
    f := BNeffatia勒法蒂亚.(*勒法蒂亚NeffatiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neffatia",
		TitleName: "勒法蒂亚",
		TitleCode: "b_neffatia",
	}
}
