package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布卡迈勒AbukamalBarony struct {
	feud.BaseBarony
}

var BAbukamal阿布卡迈勒 feud.Barony = &阿布卡迈勒AbukamalBarony{}

func init() {
    f := BAbukamal阿布卡迈勒.(*阿布卡迈勒AbukamalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abukamal",
		TitleName: "阿布卡迈勒",
		TitleCode: "b_abukamal",
	}
}
