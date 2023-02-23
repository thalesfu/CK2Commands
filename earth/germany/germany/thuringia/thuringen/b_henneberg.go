package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亨讷贝格HennebergBarony struct {
	feud.BaseBarony
}

var BHenneberg亨讷贝格 feud.Barony = &亨讷贝格HennebergBarony{}

func init() {
	f := BHenneberg亨讷贝格.(*亨讷贝格HennebergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "henneberg",
		TitleName: "亨讷贝格",
		TitleCode: "b_henneberg",
	}
}
