package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯康芒RoscommonBarony struct {
	feud.BaseBarony
}

var BRoscommon罗斯康芒 feud.Barony = &罗斯康芒RoscommonBarony{}

func init() {
    f := BRoscommon罗斯康芒.(*罗斯康芒RoscommonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roscommon",
		TitleName: "罗斯康芒",
		TitleCode: "b_roscommon",
	}
}
