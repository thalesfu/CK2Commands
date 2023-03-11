package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥霍特尼基OkhotnykyBarony struct {
	feud.BaseBarony
}

var BOkhotnyky奥霍特尼基 feud.Barony = &奥霍特尼基OkhotnykyBarony{}

func init() {
    f := BOkhotnyky奥霍特尼基.(*奥霍特尼基OkhotnykyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okhotnyky",
		TitleName: "奥霍特尼基",
		TitleCode: "b_okhotnyky",
	}
}
