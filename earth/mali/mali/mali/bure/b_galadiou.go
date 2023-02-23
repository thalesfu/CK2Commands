package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉迪乌GaladiouBarony struct {
	feud.BaseBarony
}

var BGaladiou加拉迪乌 feud.Barony = &加拉迪乌GaladiouBarony{}

func init() {
	f := BGaladiou加拉迪乌.(*加拉迪乌GaladiouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galadiou",
		TitleName: "加拉迪乌",
		TitleCode: "b_galadiou",
	}
}
