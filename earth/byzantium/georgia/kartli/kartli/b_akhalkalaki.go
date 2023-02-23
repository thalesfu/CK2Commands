package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿哈尔卡拉基AkhalkalakiBarony struct {
	feud.BaseBarony
}

var BAkhalkalaki阿哈尔卡拉基 feud.Barony = &阿哈尔卡拉基AkhalkalakiBarony{}

func init() {
	f := BAkhalkalaki阿哈尔卡拉基.(*阿哈尔卡拉基AkhalkalakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhalkalaki",
		TitleName: "阿哈尔卡拉基",
		TitleCode: "b_akhalkalaki",
	}
}
