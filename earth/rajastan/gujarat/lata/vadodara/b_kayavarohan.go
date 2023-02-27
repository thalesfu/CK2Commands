package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦耶伐卢汉拿KayavarohanBarony struct {
	feud.BaseBarony
}

var BKayavarohan迦耶伐卢汉拿 feud.Barony = &迦耶伐卢汉拿KayavarohanBarony{}

func init() {
    f := BKayavarohan迦耶伐卢汉拿.(*迦耶伐卢汉拿KayavarohanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kayavarohan",
		TitleName: "迦耶伐卢汉拿",
		TitleCode: "b_kayavarohan",
	}
}
