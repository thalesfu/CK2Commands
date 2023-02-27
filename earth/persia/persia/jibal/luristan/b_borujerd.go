package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博鲁杰尔德BorujerdBarony struct {
	feud.BaseBarony
}

var BBorujerd博鲁杰尔德 feud.Barony = &博鲁杰尔德BorujerdBarony{}

func init() {
    f := BBorujerd博鲁杰尔德.(*博鲁杰尔德BorujerdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borujerd",
		TitleName: "博鲁杰尔德",
		TitleCode: "b_borujerd",
	}
}
