package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗托内CrotoneBarony struct {
	feud.BaseBarony
}

var BCrotone克罗托内 feud.Barony = &克罗托内CrotoneBarony{}

func init() {
    f := BCrotone克罗托内.(*克罗托内CrotoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crotone",
		TitleName: "克罗托内",
		TitleCode: "b_crotone",
	}
}
