package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔齐尼UlcinjBarony struct {
	feud.BaseBarony
}

var BUlcinj乌尔齐尼 feud.Barony = &乌尔齐尼UlcinjBarony{}

func init() {
    f := BUlcinj乌尔齐尼.(*乌尔齐尼UlcinjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulcinj",
		TitleName: "乌尔齐尼",
		TitleCode: "b_ulcinj",
	}
}
