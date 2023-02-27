package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯恩斯坦BernsteinBarony struct {
	feud.BaseBarony
}

var BBernstein伯恩斯坦 feud.Barony = &伯恩斯坦BernsteinBarony{}

func init() {
    f := BBernstein伯恩斯坦.(*伯恩斯坦BernsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bernstein",
		TitleName: "伯恩斯坦",
		TitleCode: "b_bernstein",
	}
}
