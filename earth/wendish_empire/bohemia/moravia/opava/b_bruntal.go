package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布伦塔尔BruntalBarony struct {
	feud.BaseBarony
}

var BBruntal布伦塔尔 feud.Barony = &布伦塔尔BruntalBarony{}

func init() {
    f := BBruntal布伦塔尔.(*布伦塔尔BruntalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bruntal",
		TitleName: "布伦塔尔",
		TitleCode: "b_bruntal",
	}
}
