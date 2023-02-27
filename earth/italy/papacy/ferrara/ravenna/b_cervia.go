package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔维亚CerviaBarony struct {
	feud.BaseBarony
}

var BCervia切尔维亚 feud.Barony = &切尔维亚CerviaBarony{}

func init() {
    f := BCervia切尔维亚.(*切尔维亚CerviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cervia",
		TitleName: "切尔维亚",
		TitleCode: "b_cervia",
	}
}
