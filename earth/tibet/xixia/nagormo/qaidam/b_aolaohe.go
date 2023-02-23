package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嗷唠河AolaoheBarony struct {
	feud.BaseBarony
}

var BAolaohe嗷唠河 feud.Barony = &嗷唠河AolaoheBarony{}

func init() {
	f := BAolaohe嗷唠河.(*嗷唠河AolaoheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aolaohe",
		TitleName: "嗷唠河",
		TitleCode: "b_aolaohe",
	}
}
