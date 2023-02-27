package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔戈勒GorgolBarony struct {
	feud.BaseBarony
}

var BGorgol戈尔戈勒 feud.Barony = &戈尔戈勒GorgolBarony{}

func init() {
    f := BGorgol戈尔戈勒.(*戈尔戈勒GorgolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorgol",
		TitleName: "戈尔戈勒",
		TitleCode: "b_gorgol",
	}
}
