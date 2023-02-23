package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃克巴坦那EcbatanaBarony struct {
	feud.BaseBarony
}

var BEcbatana埃克巴坦那 feud.Barony = &埃克巴坦那EcbatanaBarony{}

func init() {
	f := BEcbatana埃克巴坦那.(*埃克巴坦那EcbatanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ecbatana",
		TitleName: "埃克巴坦那",
		TitleCode: "b_ecbatana",
	}
}
