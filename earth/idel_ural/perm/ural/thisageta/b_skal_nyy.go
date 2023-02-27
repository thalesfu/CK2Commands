package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡利内Skal_nyyBarony struct {
	feud.BaseBarony
}

var BSkal_nyy斯卡利内 feud.Barony = &斯卡利内Skal_nyyBarony{}

func init() {
    f := BSkal_nyy斯卡利内.(*斯卡利内Skal_nyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skal_nyy",
		TitleName: "斯卡利内",
		TitleCode: "b_skal_nyy",
	}
}
