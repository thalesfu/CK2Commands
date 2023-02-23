package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱韦恩JuwaynBarony struct {
	feud.BaseBarony
}

var BJuwayn朱韦恩 feud.Barony = &朱韦恩JuwaynBarony{}

func init() {
	f := BJuwayn朱韦恩.(*朱韦恩JuwaynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juwayn",
		TitleName: "朱韦恩",
		TitleCode: "b_juwayn",
	}
}
