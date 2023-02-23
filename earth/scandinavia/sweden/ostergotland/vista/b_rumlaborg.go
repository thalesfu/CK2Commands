package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁姆拉堡RumlaborgBarony struct {
	feud.BaseBarony
}

var BRumlaborg鲁姆拉堡 feud.Barony = &鲁姆拉堡RumlaborgBarony{}

func init() {
	f := BRumlaborg鲁姆拉堡.(*鲁姆拉堡RumlaborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rumlaborg",
		TitleName: "鲁姆拉堡",
		TitleCode: "b_rumlaborg",
	}
}
