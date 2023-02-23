package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪兹达HaozdarBarony struct {
	feud.BaseBarony
}

var BHaozdar豪兹达 feud.Barony = &豪兹达HaozdarBarony{}

func init() {
	f := BHaozdar豪兹达.(*豪兹达HaozdarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haozdar",
		TitleName: "豪兹达",
		TitleCode: "b_haozdar",
	}
}
