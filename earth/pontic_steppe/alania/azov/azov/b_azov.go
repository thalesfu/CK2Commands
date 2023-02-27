package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚速AzovBarony struct {
	feud.BaseBarony
}

var BAzov亚速 feud.Barony = &亚速AzovBarony{}

func init() {
    f := BAzov亚速.(*亚速AzovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azov",
		TitleName: "亚速",
		TitleCode: "b_azov",
	}
}
