package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季托连科TitorenkoBarony struct {
	feud.BaseBarony
}

var BTitorenko季托连科 feud.Barony = &季托连科TitorenkoBarony{}

func init() {
    f := BTitorenko季托连科.(*季托连科TitorenkoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "titorenko",
		TitleName: "季托连科",
		TitleCode: "b_titorenko",
	}
}
