package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尔戈斯PyrgosBarony struct {
	feud.BaseBarony
}

var BPyrgos皮尔戈斯 feud.Barony = &皮尔戈斯PyrgosBarony{}

func init() {
    f := BPyrgos皮尔戈斯.(*皮尔戈斯PyrgosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyrgos",
		TitleName: "皮尔戈斯",
		TitleCode: "b_pyrgos",
	}
}
