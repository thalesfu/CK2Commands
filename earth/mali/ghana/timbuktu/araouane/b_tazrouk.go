package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔兹鲁克TazroukBarony struct {
	feud.BaseBarony
}

var BTazrouk塔兹鲁克 feud.Barony = &塔兹鲁克TazroukBarony{}

func init() {
    f := BTazrouk塔兹鲁克.(*塔兹鲁克TazroukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tazrouk",
		TitleName: "塔兹鲁克",
		TitleCode: "b_tazrouk",
	}
}
