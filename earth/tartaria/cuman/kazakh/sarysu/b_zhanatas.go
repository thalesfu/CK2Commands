package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 札纳塔斯ZhanatasBarony struct {
	feud.BaseBarony
}

var BZhanatas札纳塔斯 feud.Barony = &札纳塔斯ZhanatasBarony{}

func init() {
	f := BZhanatas札纳塔斯.(*札纳塔斯ZhanatasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhanatas",
		TitleName: "札纳塔斯",
		TitleCode: "b_zhanatas",
	}
}
