package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本阿鲁斯BenarousBarony struct {
	feud.BaseBarony
}

var BBenarous本阿鲁斯 feud.Barony = &本阿鲁斯BenarousBarony{}

func init() {
    f := BBenarous本阿鲁斯.(*本阿鲁斯BenarousBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benarous",
		TitleName: "本阿鲁斯",
		TitleCode: "b_benarous",
	}
}
