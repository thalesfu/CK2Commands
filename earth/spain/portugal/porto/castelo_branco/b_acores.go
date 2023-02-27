package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚速尔AcoresBarony struct {
	feud.BaseBarony
}

var BAcores亚速尔 feud.Barony = &亚速尔AcoresBarony{}

func init() {
    f := BAcores亚速尔.(*亚速尔AcoresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acores",
		TitleName: "亚速尔",
		TitleCode: "b_acores",
	}
}
