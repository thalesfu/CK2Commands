package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷姆BouremBarony struct {
	feud.BaseBarony
}

var BBourem布雷姆 feud.Barony = &布雷姆BouremBarony{}

func init() {
    f := BBourem布雷姆.(*布雷姆BouremBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourem",
		TitleName: "布雷姆",
		TitleCode: "b_bourem",
	}
}
