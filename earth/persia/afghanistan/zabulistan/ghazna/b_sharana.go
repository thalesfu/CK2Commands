package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑兰那SharanaBarony struct {
	feud.BaseBarony
}

var BSharana娑兰那 feud.Barony = &娑兰那SharanaBarony{}

func init() {
	f := BSharana娑兰那.(*娑兰那SharanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharana",
		TitleName: "娑兰那",
		TitleCode: "b_sharana",
	}
}
