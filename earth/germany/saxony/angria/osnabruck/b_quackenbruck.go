package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夸肯布吕克QuackenbruckBarony struct {
	feud.BaseBarony
}

var BQuackenbruck夸肯布吕克 feud.Barony = &夸肯布吕克QuackenbruckBarony{}

func init() {
	f := BQuackenbruck夸肯布吕克.(*夸肯布吕克QuackenbruckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quackenbruck",
		TitleName: "夸肯布吕克",
		TitleCode: "b_quackenbruck",
	}
}
