package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安托尼耶沃_西斯基AntonievosiyskyBarony struct {
	feud.BaseBarony
}

var BAntonievosiysky安托尼耶沃_西斯基 feud.Barony = &安托尼耶沃_西斯基AntonievosiyskyBarony{}

func init() {
    f := BAntonievosiysky安托尼耶沃_西斯基.(*安托尼耶沃_西斯基AntonievosiyskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antonievosiysky",
		TitleName: "安托尼耶沃_西斯基",
		TitleCode: "b_antonievosiysky",
	}
}
