package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃时健阇HajiganjBarony struct {
	feud.BaseBarony
}

var BHajiganj诃时健阇 feud.Barony = &诃时健阇HajiganjBarony{}

func init() {
    f := BHajiganj诃时健阇.(*诃时健阇HajiganjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hajiganj",
		TitleName: "诃时健阇",
		TitleCode: "b_hajiganj",
	}
}
