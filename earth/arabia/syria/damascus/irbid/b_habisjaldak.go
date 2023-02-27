package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈比斯贾勒达克HabisjaldakBarony struct {
	feud.BaseBarony
}

var BHabisjaldak哈比斯贾勒达克 feud.Barony = &哈比斯贾勒达克HabisjaldakBarony{}

func init() {
    f := BHabisjaldak哈比斯贾勒达克.(*哈比斯贾勒达克HabisjaldakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "habisjaldak",
		TitleName: "哈比斯贾勒达克",
		TitleCode: "b_habisjaldak",
	}
}
