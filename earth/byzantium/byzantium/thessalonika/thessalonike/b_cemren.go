package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切木兰CemrenBarony struct {
	feud.BaseBarony
}

var BCemren切木兰 feud.Barony = &切木兰CemrenBarony{}

func init() {
	f := BCemren切木兰.(*切木兰CemrenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cemren",
		TitleName: "切木兰",
		TitleCode: "b_cemren",
	}
}
