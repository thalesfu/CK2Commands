package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尤瓦AyuvaBarony struct {
	feud.BaseBarony
}

var BAyuva阿尤瓦 feud.Barony = &阿尤瓦AyuvaBarony{}

func init() {
    f := BAyuva阿尤瓦.(*阿尤瓦AyuvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayuva",
		TitleName: "阿尤瓦",
		TitleCode: "b_ayuva",
	}
}
