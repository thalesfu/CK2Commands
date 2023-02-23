package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯诺亚KrasnoyaBarony struct {
	feud.BaseBarony
}

var BKrasnoya克拉斯诺亚 feud.Barony = &克拉斯诺亚KrasnoyaBarony{}

func init() {
	f := BKrasnoya克拉斯诺亚.(*克拉斯诺亚KrasnoyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasnoya",
		TitleName: "克拉斯诺亚",
		TitleCode: "b_krasnoya",
	}
}
