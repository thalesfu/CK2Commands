package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恒伽补罗GangapurBarony struct {
	feud.BaseBarony
}

var BGangapur恒伽补罗 feud.Barony = &恒伽补罗GangapurBarony{}

func init() {
	f := BGangapur恒伽补罗.(*恒伽补罗GangapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangapur",
		TitleName: "恒伽补罗",
		TitleCode: "b_gangapur",
	}
}
