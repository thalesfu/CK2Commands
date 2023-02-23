package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 金罗斯KinrossBarony struct {
	feud.BaseBarony
}

var BKinross金罗斯 feud.Barony = &金罗斯KinrossBarony{}

func init() {
	f := BKinross金罗斯.(*金罗斯KinrossBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinross",
		TitleName: "金罗斯",
		TitleCode: "b_kinross",
	}
}
