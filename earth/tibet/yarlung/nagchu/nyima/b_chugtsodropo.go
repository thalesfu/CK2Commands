package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 秋措卓波ChugtsodropoBarony struct {
	feud.BaseBarony
}

var BChugtsodropo秋措卓波 feud.Barony = &秋措卓波ChugtsodropoBarony{}

func init() {
    f := BChugtsodropo秋措卓波.(*秋措卓波ChugtsodropoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chugtsodropo",
		TitleName: "秋措卓波",
		TitleCode: "b_chugtsodropo",
	}
}
