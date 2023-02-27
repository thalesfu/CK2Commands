package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗尔FeursBarony struct {
	feud.BaseBarony
}

var BFeurs弗尔 feud.Barony = &弗尔FeursBarony{}

func init() {
    f := BFeurs弗尔.(*弗尔FeursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "feurs",
		TitleName: "弗尔",
		TitleCode: "b_feurs",
	}
}
