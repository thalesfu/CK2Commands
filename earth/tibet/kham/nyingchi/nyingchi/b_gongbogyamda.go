package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 工布江达GongbogyamdaBarony struct {
	feud.BaseBarony
}

var BGongbogyamda工布江达 feud.Barony = &工布江达GongbogyamdaBarony{}

func init() {
	f := BGongbogyamda工布江达.(*工布江达GongbogyamdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gongbogyamda",
		TitleName: "工布江达",
		TitleCode: "b_gongbogyamda",
	}
}
