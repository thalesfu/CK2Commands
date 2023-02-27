package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 仁布RinbungBarony struct {
	feud.BaseBarony
}

var BRinbung仁布 feud.Barony = &仁布RinbungBarony{}

func init() {
    f := BRinbung仁布.(*仁布RinbungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rinbung",
		TitleName: "仁布",
		TitleCode: "b_rinbung",
	}
}
