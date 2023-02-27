package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉茨GalazBarony struct {
	feud.BaseBarony
}

var BGalaz加拉茨 feud.Barony = &加拉茨GalazBarony{}

func init() {
    f := BGalaz加拉茨.(*加拉茨GalazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galaz",
		TitleName: "加拉茨",
		TitleCode: "b_galaz",
	}
}
