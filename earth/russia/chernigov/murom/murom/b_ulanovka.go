package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌拉诺夫卡UlanovkaBarony struct {
	feud.BaseBarony
}

var BUlanovka乌拉诺夫卡 feud.Barony = &乌拉诺夫卡UlanovkaBarony{}

func init() {
	f := BUlanovka乌拉诺夫卡.(*乌拉诺夫卡UlanovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulanovka",
		TitleName: "乌拉诺夫卡",
		TitleCode: "b_ulanovka",
	}
}
