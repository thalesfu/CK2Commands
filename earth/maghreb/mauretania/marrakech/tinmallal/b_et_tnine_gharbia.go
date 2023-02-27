package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尼奈加尔比耶Et_tnine_gharbiaBarony struct {
	feud.BaseBarony
}

var BEt_tnine_gharbia特尼奈加尔比耶 feud.Barony = &特尼奈加尔比耶Et_tnine_gharbiaBarony{}

func init() {
    f := BEt_tnine_gharbia特尼奈加尔比耶.(*特尼奈加尔比耶Et_tnine_gharbiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "et_tnine_gharbia",
		TitleName: "特尼奈加尔比耶",
		TitleCode: "b_et_tnine_gharbia",
	}
}
