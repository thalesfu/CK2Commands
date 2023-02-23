package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 子亭ZitingBarony struct {
	feud.BaseBarony
}

var BZiting子亭 feud.Barony = &子亭ZitingBarony{}

func init() {
	f := BZiting子亭.(*子亭ZitingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ziting",
		TitleName: "子亭",
		TitleCode: "b_ziting",
	}
}
