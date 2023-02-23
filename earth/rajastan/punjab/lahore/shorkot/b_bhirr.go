package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班尔BhirrBarony struct {
	feud.BaseBarony
}

var BBhirr班尔 feud.Barony = &班尔BhirrBarony{}

func init() {
	f := BBhirr班尔.(*班尔BhirrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhirr",
		TitleName: "班尔",
		TitleCode: "b_bhirr",
	}
}
