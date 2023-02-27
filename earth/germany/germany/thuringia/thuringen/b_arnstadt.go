package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿恩施塔特ArnstadtBarony struct {
	feud.BaseBarony
}

var BArnstadt阿恩施塔特 feud.Barony = &阿恩施塔特ArnstadtBarony{}

func init() {
    f := BArnstadt阿恩施塔特.(*阿恩施塔特ArnstadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arnstadt",
		TitleName: "阿恩施塔特",
		TitleCode: "b_arnstadt",
	}
}
