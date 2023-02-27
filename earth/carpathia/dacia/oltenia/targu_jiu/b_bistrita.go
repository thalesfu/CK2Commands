package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯特里察BistritaBarony struct {
	feud.BaseBarony
}

var BBistrita比斯特里察 feud.Barony = &比斯特里察BistritaBarony{}

func init() {
    f := BBistrita比斯特里察.(*比斯特里察BistritaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bistrita",
		TitleName: "比斯特里察",
		TitleCode: "b_bistrita",
	}
}
