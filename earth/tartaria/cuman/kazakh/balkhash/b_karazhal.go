package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉扎尔KarazhalBarony struct {
	feud.BaseBarony
}

var BKarazhal卡拉扎尔 feud.Barony = &卡拉扎尔KarazhalBarony{}

func init() {
    f := BKarazhal卡拉扎尔.(*卡拉扎尔KarazhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karazhal",
		TitleName: "卡拉扎尔",
		TitleCode: "b_karazhal",
	}
}
