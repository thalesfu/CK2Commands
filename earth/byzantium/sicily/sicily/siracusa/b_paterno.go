package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕泰尔诺PaternoBarony struct {
	feud.BaseBarony
}

var BPaterno帕泰尔诺 feud.Barony = &帕泰尔诺PaternoBarony{}

func init() {
    f := BPaterno帕泰尔诺.(*帕泰尔诺PaternoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paterno",
		TitleName: "帕泰尔诺",
		TitleCode: "b_paterno",
	}
}
