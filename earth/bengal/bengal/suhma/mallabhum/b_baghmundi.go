package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆迦蒙提BaghmundiBarony struct {
	feud.BaseBarony
}

var BBaghmundi婆迦蒙提 feud.Barony = &婆迦蒙提BaghmundiBarony{}

func init() {
    f := BBaghmundi婆迦蒙提.(*婆迦蒙提BaghmundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghmundi",
		TitleName: "婆迦蒙提",
		TitleCode: "b_baghmundi",
	}
}
