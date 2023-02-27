package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚姆基YamkiBarony struct {
	feud.BaseBarony
}

var BYamki亚姆基 feud.Barony = &亚姆基YamkiBarony{}

func init() {
    f := BYamki亚姆基.(*亚姆基YamkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yamki",
		TitleName: "亚姆基",
		TitleCode: "b_yamki",
	}
}
