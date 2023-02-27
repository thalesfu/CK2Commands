package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拉什JarashBarony struct {
	feud.BaseBarony
}

var BJarash杰拉什 feud.Barony = &杰拉什JarashBarony{}

func init() {
    f := BJarash杰拉什.(*杰拉什JarashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarash",
		TitleName: "杰拉什",
		TitleCode: "b_jarash",
	}
}
