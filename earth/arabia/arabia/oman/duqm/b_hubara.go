package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡巴拉HubaraBarony struct {
	feud.BaseBarony
}

var BHubara胡巴拉 feud.Barony = &胡巴拉HubaraBarony{}

func init() {
    f := BHubara胡巴拉.(*胡巴拉HubaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hubara",
		TitleName: "胡巴拉",
		TitleCode: "b_hubara",
	}
}
