package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托木TomBarony struct {
	feud.BaseBarony
}

var BTom托木 feud.Barony = &托木TomBarony{}

func init() {
    f := BTom托木.(*托木TomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tom",
		TitleName: "托木",
		TitleCode: "b_tom",
	}
}
