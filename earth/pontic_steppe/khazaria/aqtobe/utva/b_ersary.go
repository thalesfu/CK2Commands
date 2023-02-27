package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶尔萨雷ErsaryBarony struct {
	feud.BaseBarony
}

var BErsary叶尔萨雷 feud.Barony = &叶尔萨雷ErsaryBarony{}

func init() {
    f := BErsary叶尔萨雷.(*叶尔萨雷ErsaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ersary",
		TitleName: "叶尔萨雷",
		TitleCode: "b_ersary",
	}
}
