package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺让NogentBarony struct {
	feud.BaseBarony
}

var BNogent诺让 feud.Barony = &诺让NogentBarony{}

func init() {
	f := BNogent诺让.(*诺让NogentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nogent",
		TitleName: "诺让",
		TitleCode: "b_nogent",
	}
}
