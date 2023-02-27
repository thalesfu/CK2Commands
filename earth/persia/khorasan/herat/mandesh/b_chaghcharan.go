package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查赫查兰ChaghcharanBarony struct {
	feud.BaseBarony
}

var BChaghcharan查赫查兰 feud.Barony = &查赫查兰ChaghcharanBarony{}

func init() {
    f := BChaghcharan查赫查兰.(*查赫查兰ChaghcharanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaghcharan",
		TitleName: "查赫查兰",
		TitleCode: "b_chaghcharan",
	}
}
