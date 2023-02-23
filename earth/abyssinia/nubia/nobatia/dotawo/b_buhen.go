package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布亨BuhenBarony struct {
	feud.BaseBarony
}

var BBuhen布亨 feud.Barony = &布亨BuhenBarony{}

func init() {
	f := BBuhen布亨.(*布亨BuhenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buhen",
		TitleName: "布亨",
		TitleCode: "b_buhen",
	}
}
