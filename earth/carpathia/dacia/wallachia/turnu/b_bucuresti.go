package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布加勒斯特BucurestiBarony struct {
	feud.BaseBarony
}

var BBucuresti布加勒斯特 feud.Barony = &布加勒斯特BucurestiBarony{}

func init() {
	f := BBucuresti布加勒斯特.(*布加勒斯特BucurestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bucuresti",
		TitleName: "布加勒斯特",
		TitleCode: "b_bucuresti",
	}
}
