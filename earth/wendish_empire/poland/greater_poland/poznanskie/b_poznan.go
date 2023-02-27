package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波兹南PoznanBarony struct {
	feud.BaseBarony
}

var BPoznan波兹南 feud.Barony = &波兹南PoznanBarony{}

func init() {
    f := BPoznan波兹南.(*波兹南PoznanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poznan",
		TitleName: "波兹南",
		TitleCode: "b_poznan",
	}
}
