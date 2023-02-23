package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦达帕里VadapallyBarony struct {
	feud.BaseBarony
}

var BVadapally瓦达帕里 feud.Barony = &瓦达帕里VadapallyBarony{}

func init() {
	f := BVadapally瓦达帕里.(*瓦达帕里VadapallyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vadapally",
		TitleName: "瓦达帕里",
		TitleCode: "b_vadapally",
	}
}
