package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本比夫雷BembibreBarony struct {
	feud.BaseBarony
}

var BBembibre本比夫雷 feud.Barony = &本比夫雷BembibreBarony{}

func init() {
	f := BBembibre本比夫雷.(*本比夫雷BembibreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bembibre",
		TitleName: "本比夫雷",
		TitleCode: "b_bembibre",
	}
}
