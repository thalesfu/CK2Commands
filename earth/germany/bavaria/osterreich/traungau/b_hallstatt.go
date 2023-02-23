package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔施塔特HallstattBarony struct {
	feud.BaseBarony
}

var BHallstatt哈尔施塔特 feud.Barony = &哈尔施塔特HallstattBarony{}

func init() {
	f := BHallstatt哈尔施塔特.(*哈尔施塔特HallstattBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hallstatt",
		TitleName: "哈尔施塔特",
		TitleCode: "b_hallstatt",
	}
}
