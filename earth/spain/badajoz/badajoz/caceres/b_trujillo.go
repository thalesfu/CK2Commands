package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特鲁希略TrujilloBarony struct {
	feud.BaseBarony
}

var BTrujillo特鲁希略 feud.Barony = &特鲁希略TrujilloBarony{}

func init() {
    f := BTrujillo特鲁希略.(*特鲁希略TrujilloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trujillo",
		TitleName: "特鲁希略",
		TitleCode: "b_trujillo",
	}
}
