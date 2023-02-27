package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古林GuulinBarony struct {
	feud.BaseBarony
}

var BGuulin古林 feud.Barony = &古林GuulinBarony{}

func init() {
    f := BGuulin古林.(*古林GuulinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guulin",
		TitleName: "古林",
		TitleCode: "b_guulin",
	}
}
