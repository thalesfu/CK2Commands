package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨兰达SarandaBarony struct {
	feud.BaseBarony
}

var BSaranda萨兰达 feud.Barony = &萨兰达SarandaBarony{}

func init() {
    f := BSaranda萨兰达.(*萨兰达SarandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saranda",
		TitleName: "萨兰达",
		TitleCode: "b_saranda",
	}
}
