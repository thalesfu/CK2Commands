package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉帕尼TrapaniBarony struct {
	feud.BaseBarony
}

var BTrapani特拉帕尼 feud.Barony = &特拉帕尼TrapaniBarony{}

func init() {
    f := BTrapani特拉帕尼.(*特拉帕尼TrapaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trapani",
		TitleName: "特拉帕尼",
		TitleCode: "b_trapani",
	}
}
