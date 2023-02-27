package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾劳EylauBarony struct {
	feud.BaseBarony
}

var BEylau艾劳 feud.Barony = &艾劳EylauBarony{}

func init() {
    f := BEylau艾劳.(*艾劳EylauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eylau",
		TitleName: "艾劳",
		TitleCode: "b_eylau",
	}
}
