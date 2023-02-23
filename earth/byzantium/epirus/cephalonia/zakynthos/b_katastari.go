package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔斯塔里KatastariBarony struct {
	feud.BaseBarony
}

var BKatastari卡塔斯塔里 feud.Barony = &卡塔斯塔里KatastariBarony{}

func init() {
	f := BKatastari卡塔斯塔里.(*卡塔斯塔里KatastariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katastari",
		TitleName: "卡塔斯塔里",
		TitleCode: "b_katastari",
	}
}
