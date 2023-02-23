package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕伦西亚PalenciaBarony struct {
	feud.BaseBarony
}

var BPalencia帕伦西亚 feud.Barony = &帕伦西亚PalenciaBarony{}

func init() {
	f := BPalencia帕伦西亚.(*帕伦西亚PalenciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palencia",
		TitleName: "帕伦西亚",
		TitleCode: "b_palencia",
	}
}
