package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尼图瓦KnituvaBarony struct {
	feud.BaseBarony
}

var BKnituva克尼图瓦 feud.Barony = &克尼图瓦KnituvaBarony{}

func init() {
    f := BKnituva克尼图瓦.(*克尼图瓦KnituvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knituva",
		TitleName: "克尼图瓦",
		TitleCode: "b_knituva",
	}
}
