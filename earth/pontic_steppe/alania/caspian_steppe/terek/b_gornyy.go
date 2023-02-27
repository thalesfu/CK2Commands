package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔内GornyyBarony struct {
	feud.BaseBarony
}

var BGornyy戈尔内 feud.Barony = &戈尔内GornyyBarony{}

func init() {
    f := BGornyy戈尔内.(*戈尔内GornyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gornyy",
		TitleName: "戈尔内",
		TitleCode: "b_gornyy",
	}
}
