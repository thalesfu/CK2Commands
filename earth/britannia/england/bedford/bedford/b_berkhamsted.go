package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯克姆斯特德BerkhamstedBarony struct {
	feud.BaseBarony
}

var BBerkhamsted伯克姆斯特德 feud.Barony = &伯克姆斯特德BerkhamstedBarony{}

func init() {
	f := BBerkhamsted伯克姆斯特德.(*伯克姆斯特德BerkhamstedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berkhamsted",
		TitleName: "伯克姆斯特德",
		TitleCode: "b_berkhamsted",
	}
}
