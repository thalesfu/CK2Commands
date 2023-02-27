package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西伯库伦Shivee_khurenBarony struct {
	feud.BaseBarony
}

var BShivee_khuren西伯库伦 feud.Barony = &西伯库伦Shivee_khurenBarony{}

func init() {
    f := BShivee_khuren西伯库伦.(*西伯库伦Shivee_khurenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shivee_khuren",
		TitleName: "西伯库伦",
		TitleCode: "b_shivee_khuren",
	}
}
