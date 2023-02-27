package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡德穆斯AlqadmusBarony struct {
	feud.BaseBarony
}

var BAlqadmus卡德穆斯 feud.Barony = &卡德穆斯AlqadmusBarony{}

func init() {
    f := BAlqadmus卡德穆斯.(*卡德穆斯AlqadmusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqadmus",
		TitleName: "卡德穆斯",
		TitleCode: "b_alqadmus",
	}
}
