package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢那瓦陀LunawadaBarony struct {
	feud.BaseBarony
}

var BLunawada卢那瓦陀 feud.Barony = &卢那瓦陀LunawadaBarony{}

func init() {
    f := BLunawada卢那瓦陀.(*卢那瓦陀LunawadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lunawada",
		TitleName: "卢那瓦陀",
		TitleCode: "b_lunawada",
	}
}
