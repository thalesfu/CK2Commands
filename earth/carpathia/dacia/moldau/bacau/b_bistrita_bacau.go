package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯特里察Bistrita_bacauBarony struct {
	feud.BaseBarony
}

var BBistrita_bacau比斯特里察 feud.Barony = &比斯特里察Bistrita_bacauBarony{}

func init() {
    f := BBistrita_bacau比斯特里察.(*比斯特里察Bistrita_bacauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bistrita_bacau",
		TitleName: "比斯特里察",
		TitleCode: "b_bistrita_bacau",
	}
}
