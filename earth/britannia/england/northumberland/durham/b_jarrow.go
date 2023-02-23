package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾罗JarrowBarony struct {
	feud.BaseBarony
}

var BJarrow贾罗 feud.Barony = &贾罗JarrowBarony{}

func init() {
	f := BJarrow贾罗.(*贾罗JarrowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarrow",
		TitleName: "贾罗",
		TitleCode: "b_jarrow",
	}
}
