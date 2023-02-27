package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞拉莱SalalaBarony struct {
	feud.BaseBarony
}

var BSalala塞拉莱 feud.Barony = &塞拉莱SalalaBarony{}

func init() {
    f := BSalala塞拉莱.(*塞拉莱SalalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salala",
		TitleName: "塞拉莱",
		TitleCode: "b_salala",
	}
}
