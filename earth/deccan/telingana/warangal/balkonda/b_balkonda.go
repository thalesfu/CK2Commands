package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗军荼BalkondaBarony struct {
	feud.BaseBarony
}

var BBalkonda婆罗军荼 feud.Barony = &婆罗军荼BalkondaBarony{}

func init() {
	f := BBalkonda婆罗军荼.(*婆罗军荼BalkondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balkonda",
		TitleName: "婆罗军荼",
		TitleCode: "b_balkonda",
	}
}
