package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博斯塔德BaastadBarony struct {
	feud.BaseBarony
}

var BBaastad博斯塔德 feud.Barony = &博斯塔德BaastadBarony{}

func init() {
    f := BBaastad博斯塔德.(*博斯塔德BaastadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baastad",
		TitleName: "博斯塔德",
		TitleCode: "b_baastad",
	}
}
