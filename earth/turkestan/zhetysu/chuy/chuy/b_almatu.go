package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉木图AlmatuBarony struct {
	feud.BaseBarony
}

var BAlmatu阿拉木图 feud.Barony = &阿拉木图AlmatuBarony{}

func init() {
    f := BAlmatu阿拉木图.(*阿拉木图AlmatuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almatu",
		TitleName: "阿拉木图",
		TitleCode: "b_almatu",
	}
}
