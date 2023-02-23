package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔巴哈台TarbagataiBarony struct {
	feud.BaseBarony
}

var BTarbagatai塔尔巴哈台 feud.Barony = &塔尔巴哈台TarbagataiBarony{}

func init() {
	f := BTarbagatai塔尔巴哈台.(*塔尔巴哈台TarbagataiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarbagatai",
		TitleName: "塔尔巴哈台",
		TitleCode: "b_tarbagatai",
	}
}
