package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫拉西亚卜AfrasiyabBarony struct {
	feud.BaseBarony
}

var BAfrasiyab阿夫拉西亚卜 feud.Barony = &阿夫拉西亚卜AfrasiyabBarony{}

func init() {
    f := BAfrasiyab阿夫拉西亚卜.(*阿夫拉西亚卜AfrasiyabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afrasiyab",
		TitleName: "阿夫拉西亚卜",
		TitleCode: "b_afrasiyab",
	}
}
