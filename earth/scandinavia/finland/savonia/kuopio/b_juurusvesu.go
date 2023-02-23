package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤鲁斯湖JuurusvesuBarony struct {
	feud.BaseBarony
}

var BJuurusvesu尤鲁斯湖 feud.Barony = &尤鲁斯湖JuurusvesuBarony{}

func init() {
	f := BJuurusvesu尤鲁斯湖.(*尤鲁斯湖JuurusvesuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juurusvesu",
		TitleName: "尤鲁斯湖",
		TitleCode: "b_juurusvesu",
	}
}
