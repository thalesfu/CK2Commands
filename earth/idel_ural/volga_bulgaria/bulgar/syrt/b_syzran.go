package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞兹兰SyzranBarony struct {
	feud.BaseBarony
}

var BSyzran塞兹兰 feud.Barony = &塞兹兰SyzranBarony{}

func init() {
    f := BSyzran塞兹兰.(*塞兹兰SyzranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syzran",
		TitleName: "塞兹兰",
		TitleCode: "b_syzran",
	}
}
