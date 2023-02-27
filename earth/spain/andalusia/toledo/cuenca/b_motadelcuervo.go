package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫塔德尔奎尔沃MotadelcuervoBarony struct {
	feud.BaseBarony
}

var BMotadelcuervo莫塔德尔奎尔沃 feud.Barony = &莫塔德尔奎尔沃MotadelcuervoBarony{}

func init() {
    f := BMotadelcuervo莫塔德尔奎尔沃.(*莫塔德尔奎尔沃MotadelcuervoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "motadelcuervo",
		TitleName: "莫塔德尔奎尔沃",
		TitleCode: "b_motadelcuervo",
	}
}
