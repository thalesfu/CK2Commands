package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优陀耶耆厘UdayagiriBarony struct {
	feud.BaseBarony
}

var BUdayagiri优陀耶耆厘 feud.Barony = &优陀耶耆厘UdayagiriBarony{}

func init() {
    f := BUdayagiri优陀耶耆厘.(*优陀耶耆厘UdayagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udayagiri",
		TitleName: "优陀耶耆厘",
		TitleCode: "b_udayagiri",
	}
}
