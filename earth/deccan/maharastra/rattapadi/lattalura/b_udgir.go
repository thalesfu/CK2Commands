package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优陀耆厘UdgirBarony struct {
	feud.BaseBarony
}

var BUdgir优陀耆厘 feud.Barony = &优陀耆厘UdgirBarony{}

func init() {
    f := BUdgir优陀耆厘.(*优陀耆厘UdgirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udgir",
		TitleName: "优陀耆厘",
		TitleCode: "b_udgir",
	}
}
