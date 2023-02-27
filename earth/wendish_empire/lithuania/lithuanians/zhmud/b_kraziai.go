package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉日艾KraziaiBarony struct {
	feud.BaseBarony
}

var BKraziai克拉日艾 feud.Barony = &克拉日艾KraziaiBarony{}

func init() {
    f := BKraziai克拉日艾.(*克拉日艾KraziaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kraziai",
		TitleName: "克拉日艾",
		TitleCode: "b_kraziai",
	}
}
