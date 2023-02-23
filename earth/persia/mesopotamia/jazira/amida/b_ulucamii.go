package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌鲁贾米UlucamiiBarony struct {
	feud.BaseBarony
}

var BUlucamii乌鲁贾米 feud.Barony = &乌鲁贾米UlucamiiBarony{}

func init() {
	f := BUlucamii乌鲁贾米.(*乌鲁贾米UlucamiiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulucamii",
		TitleName: "乌鲁贾米",
		TitleCode: "b_ulucamii",
	}
}
