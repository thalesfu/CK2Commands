package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 支迦罗傥那ChikalthanaBarony struct {
	feud.BaseBarony
}

var BChikalthana支迦罗傥那 feud.Barony = &支迦罗傥那ChikalthanaBarony{}

func init() {
    f := BChikalthana支迦罗傥那.(*支迦罗傥那ChikalthanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chikalthana",
		TitleName: "支迦罗傥那",
		TitleCode: "b_chikalthana",
	}
}
