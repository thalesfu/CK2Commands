package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜卡洛沃BaykalovoBarony struct {
	feud.BaseBarony
}

var BBaykalovo拜卡洛沃 feud.Barony = &拜卡洛沃BaykalovoBarony{}

func init() {
    f := BBaykalovo拜卡洛沃.(*拜卡洛沃BaykalovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baykalovo",
		TitleName: "拜卡洛沃",
		TitleCode: "b_baykalovo",
	}
}
