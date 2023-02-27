package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优陀耶伐罗UdayavaraBarony struct {
	feud.BaseBarony
}

var BUdayavara优陀耶伐罗 feud.Barony = &优陀耶伐罗UdayavaraBarony{}

func init() {
    f := BUdayavara优陀耶伐罗.(*优陀耶伐罗UdayavaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udayavara",
		TitleName: "优陀耶伐罗",
		TitleCode: "b_udayavara",
	}
}
