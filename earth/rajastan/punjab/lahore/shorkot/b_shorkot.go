package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍科特ShorkotBarony struct {
	feud.BaseBarony
}

var BShorkot绍科特 feud.Barony = &绍科特ShorkotBarony{}

func init() {
	f := BShorkot绍科特.(*绍科特ShorkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shorkot",
		TitleName: "绍科特",
		TitleCode: "b_shorkot",
	}
}
