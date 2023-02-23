package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔恰托夫KurchatovBarony struct {
	feud.BaseBarony
}

var BKurchatov库尔恰托夫 feud.Barony = &库尔恰托夫KurchatovBarony{}

func init() {
	f := BKurchatov库尔恰托夫.(*库尔恰托夫KurchatovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurchatov",
		TitleName: "库尔恰托夫",
		TitleCode: "b_kurchatov",
	}
}
