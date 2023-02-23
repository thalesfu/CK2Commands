package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃里切EriceBarony struct {
	feud.BaseBarony
}

var BErice埃里切 feud.Barony = &埃里切EriceBarony{}

func init() {
	f := BErice埃里切.(*埃里切EriceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erice",
		TitleName: "埃里切",
		TitleCode: "b_erice",
	}
}
