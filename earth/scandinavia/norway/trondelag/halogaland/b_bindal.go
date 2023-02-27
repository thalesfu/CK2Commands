package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宾达BindalBarony struct {
	feud.BaseBarony
}

var BBindal宾达 feud.Barony = &宾达BindalBarony{}

func init() {
    f := BBindal宾达.(*宾达BindalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bindal",
		TitleName: "宾达",
		TitleCode: "b_bindal",
	}
}
