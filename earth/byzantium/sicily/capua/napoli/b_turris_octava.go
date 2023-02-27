package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图里斯_奥克塔瓦Turris_octavaBarony struct {
	feud.BaseBarony
}

var BTurris_octava图里斯_奥克塔瓦 feud.Barony = &图里斯_奥克塔瓦Turris_octavaBarony{}

func init() {
    f := BTurris_octava图里斯_奥克塔瓦.(*图里斯_奥克塔瓦Turris_octavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turris_octava",
		TitleName: "图里斯_奥克塔瓦",
		TitleCode: "b_turris_octava",
	}
}
