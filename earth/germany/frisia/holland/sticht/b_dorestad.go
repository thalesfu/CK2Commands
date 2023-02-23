package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多雷斯塔德DorestadBarony struct {
	feud.BaseBarony
}

var BDorestad多雷斯塔德 feud.Barony = &多雷斯塔德DorestadBarony{}

func init() {
	f := BDorestad多雷斯塔德.(*多雷斯塔德DorestadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorestad",
		TitleName: "多雷斯塔德",
		TitleCode: "b_dorestad",
	}
}
