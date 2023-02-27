package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎帕拉GampalaBarony struct {
	feud.BaseBarony
}

var BGampala坎帕拉 feud.Barony = &坎帕拉GampalaBarony{}

func init() {
    f := BGampala坎帕拉.(*坎帕拉GampalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gampala",
		TitleName: "坎帕拉",
		TitleCode: "b_gampala",
	}
}
