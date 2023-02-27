package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马克萨季哈MaksatikhaBarony struct {
	feud.BaseBarony
}

var BMaksatikha马克萨季哈 feud.Barony = &马克萨季哈MaksatikhaBarony{}

func init() {
    f := BMaksatikha马克萨季哈.(*马克萨季哈MaksatikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maksatikha",
		TitleName: "马克萨季哈",
		TitleCode: "b_maksatikha",
	}
}
