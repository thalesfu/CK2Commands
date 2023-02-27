package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰尔迪奈JardinahBarony struct {
	feud.BaseBarony
}

var BJardinah杰尔迪奈 feud.Barony = &杰尔迪奈JardinahBarony{}

func init() {
    f := BJardinah杰尔迪奈.(*杰尔迪奈JardinahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jardinah",
		TitleName: "杰尔迪奈",
		TitleCode: "b_jardinah",
	}
}
