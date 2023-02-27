package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒讷LohnBarony struct {
	feud.BaseBarony
}

var BLohn勒讷 feud.Barony = &勒讷LohnBarony{}

func init() {
    f := BLohn勒讷.(*勒讷LohnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lohn",
		TitleName: "勒讷",
		TitleCode: "b_lohn",
	}
}
