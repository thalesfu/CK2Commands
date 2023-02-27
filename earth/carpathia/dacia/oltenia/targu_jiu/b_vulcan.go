package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武尔坎VulcanBarony struct {
	feud.BaseBarony
}

var BVulcan武尔坎 feud.Barony = &武尔坎VulcanBarony{}

func init() {
    f := BVulcan武尔坎.(*武尔坎VulcanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vulcan",
		TitleName: "武尔坎",
		TitleCode: "b_vulcan",
	}
}
