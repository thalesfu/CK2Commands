package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米努达什特MinudashtBarony struct {
	feud.BaseBarony
}

var BMinudasht米努达什特 feud.Barony = &米努达什特MinudashtBarony{}

func init() {
    f := BMinudasht米努达什特.(*米努达什特MinudashtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minudasht",
		TitleName: "米努达什特",
		TitleCode: "b_minudasht",
	}
}
