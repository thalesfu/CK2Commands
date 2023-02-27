package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 畔度伽荼BandhugadhaBarony struct {
	feud.BaseBarony
}

var BBandhugadha畔度伽荼 feud.Barony = &畔度伽荼BandhugadhaBarony{}

func init() {
    f := BBandhugadha畔度伽荼.(*畔度伽荼BandhugadhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandhugadha",
		TitleName: "畔度伽荼",
		TitleCode: "b_bandhugadha",
	}
}
