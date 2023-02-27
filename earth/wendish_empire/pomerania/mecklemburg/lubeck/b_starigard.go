package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔里加德StarigardBarony struct {
	feud.BaseBarony
}

var BStarigard斯塔里加德 feud.Barony = &斯塔里加德StarigardBarony{}

func init() {
    f := BStarigard斯塔里加德.(*斯塔里加德StarigardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "starigard",
		TitleName: "斯塔里加德",
		TitleCode: "b_starigard",
	}
}
