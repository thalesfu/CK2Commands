package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇维泰拉韦斯科沃Civitella_del_vescovoBarony struct {
	feud.BaseBarony
}

var BCivitella_del_vescovo奇维泰拉韦斯科沃 feud.Barony = &奇维泰拉韦斯科沃Civitella_del_vescovoBarony{}

func init() {
    f := BCivitella_del_vescovo奇维泰拉韦斯科沃.(*奇维泰拉韦斯科沃Civitella_del_vescovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "civitella_del_vescovo",
		TitleName: "奇维泰拉韦斯科沃",
		TitleCode: "b_civitella_del_vescovo",
	}
}
