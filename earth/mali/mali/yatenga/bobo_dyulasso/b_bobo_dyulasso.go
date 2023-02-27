package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博博迪乌拉索Bobo_dyulassoBarony struct {
	feud.BaseBarony
}

var BBobo_dyulasso博博迪乌拉索 feud.Barony = &博博迪乌拉索Bobo_dyulassoBarony{}

func init() {
    f := BBobo_dyulasso博博迪乌拉索.(*博博迪乌拉索Bobo_dyulassoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bobo_dyulasso",
		TitleName: "博博迪乌拉索",
		TitleCode: "b_bobo_dyulasso",
	}
}
