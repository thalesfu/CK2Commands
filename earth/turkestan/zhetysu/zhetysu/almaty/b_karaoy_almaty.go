package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉奥伊Karaoy_almatyBarony struct {
	feud.BaseBarony
}

var BKaraoy_almaty卡拉奥伊 feud.Barony = &卡拉奥伊Karaoy_almatyBarony{}

func init() {
    f := BKaraoy_almaty卡拉奥伊.(*卡拉奥伊Karaoy_almatyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karaoy_almaty",
		TitleName: "卡拉奥伊",
		TitleCode: "b_karaoy_almaty",
	}
}
