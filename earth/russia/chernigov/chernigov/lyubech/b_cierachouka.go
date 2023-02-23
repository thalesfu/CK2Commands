package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切拉霍卡CierachoukaBarony struct {
	feud.BaseBarony
}

var BCierachouka切拉霍卡 feud.Barony = &切拉霍卡CierachoukaBarony{}

func init() {
	f := BCierachouka切拉霍卡.(*切拉霍卡CierachoukaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cierachouka",
		TitleName: "切拉霍卡",
		TitleCode: "b_cierachouka",
	}
}
