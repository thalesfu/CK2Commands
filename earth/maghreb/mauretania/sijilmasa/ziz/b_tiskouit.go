package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂斯库伊特TiskouitBarony struct {
	feud.BaseBarony
}

var BTiskouit蒂斯库伊特 feud.Barony = &蒂斯库伊特TiskouitBarony{}

func init() {
	f := BTiskouit蒂斯库伊特.(*蒂斯库伊特TiskouitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiskouit",
		TitleName: "蒂斯库伊特",
		TitleCode: "b_tiskouit",
	}
}
