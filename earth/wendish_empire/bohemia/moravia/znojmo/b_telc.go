package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔奇TelcBarony struct {
	feud.BaseBarony
}

var BTelc泰尔奇 feud.Barony = &泰尔奇TelcBarony{}

func init() {
    f := BTelc泰尔奇.(*泰尔奇TelcBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telc",
		TitleName: "泰尔奇",
		TitleCode: "b_telc",
	}
}
