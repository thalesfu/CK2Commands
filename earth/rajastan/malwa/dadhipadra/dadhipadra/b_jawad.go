package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奢瓦德JawadBarony struct {
	feud.BaseBarony
}

var BJawad奢瓦德 feud.Barony = &奢瓦德JawadBarony{}

func init() {
	f := BJawad奢瓦德.(*奢瓦德JawadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jawad",
		TitleName: "奢瓦德",
		TitleCode: "b_jawad",
	}
}
