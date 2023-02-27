package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒洛克勒LalocleBarony struct {
	feud.BaseBarony
}

var BLalocle勒洛克勒 feud.Barony = &勒洛克勒LalocleBarony{}

func init() {
    f := BLalocle勒洛克勒.(*勒洛克勒LalocleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lalocle",
		TitleName: "勒洛克勒",
		TitleCode: "b_lalocle",
	}
}
