package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠诃怛罗BehtarBarony struct {
	feud.BaseBarony
}

var BBehtar吠诃怛罗 feud.Barony = &吠诃怛罗BehtarBarony{}

func init() {
	f := BBehtar吠诃怛罗.(*吠诃怛罗BehtarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "behtar",
		TitleName: "吠诃怛罗",
		TitleCode: "b_behtar",
	}
}
