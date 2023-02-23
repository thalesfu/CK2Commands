package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩鲁贾PerugiaBarony struct {
	feud.BaseBarony
}

var BPerugia佩鲁贾 feud.Barony = &佩鲁贾PerugiaBarony{}

func init() {
	f := BPerugia佩鲁贾.(*佩鲁贾PerugiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perugia",
		TitleName: "佩鲁贾",
		TitleCode: "b_perugia",
	}
}
