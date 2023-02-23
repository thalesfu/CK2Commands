package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞卜拉泰SabratahBarony struct {
	feud.BaseBarony
}

var BSabratah塞卜拉泰 feud.Barony = &塞卜拉泰SabratahBarony{}

func init() {
	f := BSabratah塞卜拉泰.(*塞卜拉泰SabratahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabratah",
		TitleName: "塞卜拉泰",
		TitleCode: "b_sabratah",
	}
}
