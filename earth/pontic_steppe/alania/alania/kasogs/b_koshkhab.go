package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科舍尔哈布利KoshkhabBarony struct {
	feud.BaseBarony
}

var BKoshkhab科舍尔哈布利 feud.Barony = &科舍尔哈布利KoshkhabBarony{}

func init() {
    f := BKoshkhab科舍尔哈布利.(*科舍尔哈布利KoshkhabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koshkhab",
		TitleName: "科舍尔哈布利",
		TitleCode: "b_koshkhab",
	}
}
