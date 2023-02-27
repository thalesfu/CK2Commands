package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎林沙赫尔ZarrinshahrBarony struct {
	feud.BaseBarony
}

var BZarrinshahr扎林沙赫尔 feud.Barony = &扎林沙赫尔ZarrinshahrBarony{}

func init() {
    f := BZarrinshahr扎林沙赫尔.(*扎林沙赫尔ZarrinshahrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarrinshahr",
		TitleName: "扎林沙赫尔",
		TitleCode: "b_zarrinshahr",
	}
}
