package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡尔特斯塔德HulterstadBarony struct {
	feud.BaseBarony
}

var BHulterstad胡尔特斯塔德 feud.Barony = &胡尔特斯塔德HulterstadBarony{}

func init() {
	f := BHulterstad胡尔特斯塔德.(*胡尔特斯塔德HulterstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hulterstad",
		TitleName: "胡尔特斯塔德",
		TitleCode: "b_hulterstad",
	}
}
