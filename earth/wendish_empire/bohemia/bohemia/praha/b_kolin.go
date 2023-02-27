package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科林KolinBarony struct {
	feud.BaseBarony
}

var BKolin科林 feud.Barony = &科林KolinBarony{}

func init() {
    f := BKolin科林.(*科林KolinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolin",
		TitleName: "科林",
		TitleCode: "b_kolin",
	}
}
