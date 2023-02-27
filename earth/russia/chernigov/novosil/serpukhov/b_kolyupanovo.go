package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科柳帕诺沃KolyupanovoBarony struct {
	feud.BaseBarony
}

var BKolyupanovo科柳帕诺沃 feud.Barony = &科柳帕诺沃KolyupanovoBarony{}

func init() {
    f := BKolyupanovo科柳帕诺沃.(*科柳帕诺沃KolyupanovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolyupanovo",
		TitleName: "科柳帕诺沃",
		TitleCode: "b_kolyupanovo",
	}
}
