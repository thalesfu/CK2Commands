package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝哈尔SalbejarBarony struct {
	feud.BaseBarony
}

var BSalbejar贝哈尔 feud.Barony = &贝哈尔SalbejarBarony{}

func init() {
    f := BSalbejar贝哈尔.(*贝哈尔SalbejarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salbejar",
		TitleName: "贝哈尔",
		TitleCode: "b_salbejar",
	}
}
