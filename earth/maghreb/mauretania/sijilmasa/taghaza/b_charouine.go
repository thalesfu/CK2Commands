package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔温CharouineBarony struct {
	feud.BaseBarony
}

var BCharouine沙尔温 feud.Barony = &沙尔温CharouineBarony{}

func init() {
    f := BCharouine沙尔温.(*沙尔温CharouineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charouine",
		TitleName: "沙尔温",
		TitleCode: "b_charouine",
	}
}
