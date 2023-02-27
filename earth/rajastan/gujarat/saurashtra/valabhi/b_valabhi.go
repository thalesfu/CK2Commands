package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐腊毗ValabhiBarony struct {
	feud.BaseBarony
}

var BValabhi伐腊毗 feud.Barony = &伐腊毗ValabhiBarony{}

func init() {
    f := BValabhi伐腊毗.(*伐腊毗ValabhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valabhi",
		TitleName: "伐腊毗",
		TitleCode: "b_valabhi",
	}
}
