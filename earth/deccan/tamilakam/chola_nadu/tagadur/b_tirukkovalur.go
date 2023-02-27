package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提卢伽婆卢TirukkovalurBarony struct {
	feud.BaseBarony
}

var BTirukkovalur提卢伽婆卢 feud.Barony = &提卢伽婆卢TirukkovalurBarony{}

func init() {
    f := BTirukkovalur提卢伽婆卢.(*提卢伽婆卢TirukkovalurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirukkovalur",
		TitleName: "提卢伽婆卢",
		TitleCode: "b_tirukkovalur",
	}
}
