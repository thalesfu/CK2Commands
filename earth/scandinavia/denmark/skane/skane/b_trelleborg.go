package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷勒堡TrelleborgBarony struct {
	feud.BaseBarony
}

var BTrelleborg特雷勒堡 feud.Barony = &特雷勒堡TrelleborgBarony{}

func init() {
    f := BTrelleborg特雷勒堡.(*特雷勒堡TrelleborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trelleborg",
		TitleName: "特雷勒堡",
		TitleCode: "b_trelleborg",
	}
}
