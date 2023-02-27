package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁家寨DingjiazhaiBarony struct {
	feud.BaseBarony
}

var BDingjiazhai丁家寨 feud.Barony = &丁家寨DingjiazhaiBarony{}

func init() {
    f := BDingjiazhai丁家寨.(*丁家寨DingjiazhaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dingjiazhai",
		TitleName: "丁家寨",
		TitleCode: "b_dingjiazhai",
	}
}
