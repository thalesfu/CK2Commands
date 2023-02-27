package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉希德RosettaBarony struct {
	feud.BaseBarony
}

var BRosetta拉希德 feud.Barony = &拉希德RosettaBarony{}

func init() {
    f := BRosetta拉希德.(*拉希德RosettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rosetta",
		TitleName: "拉希德",
		TitleCode: "b_rosetta",
	}
}
