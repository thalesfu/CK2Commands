package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰塔温TataouineBarony struct {
	feud.BaseBarony
}

var BTataouine泰塔温 feud.Barony = &泰塔温TataouineBarony{}

func init() {
    f := BTataouine泰塔温.(*泰塔温TataouineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tataouine",
		TitleName: "泰塔温",
		TitleCode: "b_tataouine",
	}
}
