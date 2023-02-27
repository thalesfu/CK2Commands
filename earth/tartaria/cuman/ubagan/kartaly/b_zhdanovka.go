package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日丹诺夫卡ZhdanovkaBarony struct {
	feud.BaseBarony
}

var BZhdanovka日丹诺夫卡 feud.Barony = &日丹诺夫卡ZhdanovkaBarony{}

func init() {
    f := BZhdanovka日丹诺夫卡.(*日丹诺夫卡ZhdanovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhdanovka",
		TitleName: "日丹诺夫卡",
		TitleCode: "b_zhdanovka",
	}
}
