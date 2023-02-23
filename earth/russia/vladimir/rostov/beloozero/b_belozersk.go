package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别洛奥泽罗BelozerskBarony struct {
	feud.BaseBarony
}

var BBelozersk别洛奥泽罗 feud.Barony = &别洛奥泽罗BelozerskBarony{}

func init() {
	f := BBelozersk别洛奥泽罗.(*别洛奥泽罗BelozerskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belozersk",
		TitleName: "别洛奥泽罗",
		TitleCode: "b_belozersk",
	}
}
