package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦古洛WaguloBarony struct {
	feud.BaseBarony
}

var BWagulo瓦古洛 feud.Barony = &瓦古洛WaguloBarony{}

func init() {
    f := BWagulo瓦古洛.(*瓦古洛WaguloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wagulo",
		TitleName: "瓦古洛",
		TitleCode: "b_wagulo",
	}
}
