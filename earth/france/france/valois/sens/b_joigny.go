package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹瓦尼JoignyBarony struct {
	feud.BaseBarony
}

var BJoigny茹瓦尼 feud.Barony = &茹瓦尼JoignyBarony{}

func init() {
    f := BJoigny茹瓦尼.(*茹瓦尼JoignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joigny",
		TitleName: "茹瓦尼",
		TitleCode: "b_joigny",
	}
}
