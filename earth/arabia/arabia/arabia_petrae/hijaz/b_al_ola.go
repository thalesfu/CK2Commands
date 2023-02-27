package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧拉Al_olaBarony struct {
	feud.BaseBarony
}

var BAl_ola欧拉 feud.Barony = &欧拉Al_olaBarony{}

func init() {
    f := BAl_ola欧拉.(*欧拉Al_olaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_ola",
		TitleName: "欧拉",
		TitleCode: "b_al_ola",
	}
}
