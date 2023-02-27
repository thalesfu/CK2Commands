package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒尔楚克ShorshuqBarony struct {
	feud.BaseBarony
}

var BShorshuq舒尔楚克 feud.Barony = &舒尔楚克ShorshuqBarony{}

func init() {
    f := BShorshuq舒尔楚克.(*舒尔楚克ShorshuqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shorshuq",
		TitleName: "舒尔楚克",
		TitleCode: "b_shorshuq",
	}
}
