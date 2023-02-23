package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔施塔特MalstattBarony struct {
	feud.BaseBarony
}

var BMalstatt马尔施塔特 feud.Barony = &马尔施塔特MalstattBarony{}

func init() {
	f := BMalstatt马尔施塔特.(*马尔施塔特MalstattBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malstatt",
		TitleName: "马尔施塔特",
		TitleCode: "b_malstatt",
	}
}
