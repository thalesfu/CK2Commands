package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫塔尔巴赫塔BakhtarakhtarBarony struct {
	feud.BaseBarony
}

var BBakhtarakhtar阿赫塔尔巴赫塔 feud.Barony = &阿赫塔尔巴赫塔BakhtarakhtarBarony{}

func init() {
    f := BBakhtarakhtar阿赫塔尔巴赫塔.(*阿赫塔尔巴赫塔BakhtarakhtarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakhtarakhtar",
		TitleName: "阿赫塔尔巴赫塔",
		TitleCode: "b_bakhtarakhtar",
	}
}
