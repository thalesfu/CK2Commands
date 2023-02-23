package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕米耶PamiersBarony struct {
	feud.BaseBarony
}

var BPamiers帕米耶 feud.Barony = &帕米耶PamiersBarony{}

func init() {
	f := BPamiers帕米耶.(*帕米耶PamiersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pamiers",
		TitleName: "帕米耶",
		TitleCode: "b_pamiers",
	}
}
