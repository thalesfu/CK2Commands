package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆欧云QumelaiounBarony struct {
	feud.BaseBarony
}

var BQumelaioun乌姆欧云 feud.Barony = &乌姆欧云QumelaiounBarony{}

func init() {
    f := BQumelaioun乌姆欧云.(*乌姆欧云QumelaiounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qumelaioun",
		TitleName: "乌姆欧云",
		TitleCode: "b_qumelaioun",
	}
}
