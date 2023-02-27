package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌塞拉Al_usaylaBarony struct {
	feud.BaseBarony
}

var BAl_usayla乌塞拉 feud.Barony = &乌塞拉Al_usaylaBarony{}

func init() {
    f := BAl_usayla乌塞拉.(*乌塞拉Al_usaylaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_usayla",
		TitleName: "乌塞拉",
		TitleCode: "b_al_usayla",
	}
}
