package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔比CherbiBarony struct {
	feud.BaseBarony
}

var BCherbi切尔比 feud.Barony = &切尔比CherbiBarony{}

func init() {
    f := BCherbi切尔比.(*切尔比CherbiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherbi",
		TitleName: "切尔比",
		TitleCode: "b_cherbi",
	}
}
