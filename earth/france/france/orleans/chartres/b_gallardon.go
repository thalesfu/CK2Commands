package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉尔东GallardonBarony struct {
	feud.BaseBarony
}

var BGallardon加拉尔东 feud.Barony = &加拉尔东GallardonBarony{}

func init() {
    f := BGallardon加拉尔东.(*加拉尔东GallardonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gallardon",
		TitleName: "加拉尔东",
		TitleCode: "b_gallardon",
	}
}
