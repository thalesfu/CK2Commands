package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚瓦特马尔YavatmalBarony struct {
	feud.BaseBarony
}

var BYavatmal亚瓦特马尔 feud.Barony = &亚瓦特马尔YavatmalBarony{}

func init() {
	f := BYavatmal亚瓦特马尔.(*亚瓦特马尔YavatmalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yavatmal",
		TitleName: "亚瓦特马尔",
		TitleCode: "b_yavatmal",
	}
}
