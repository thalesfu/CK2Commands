package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴什库尔干BashqorghanBarony struct {
	feud.BaseBarony
}

var BBashqorghan巴什库尔干 feud.Barony = &巴什库尔干BashqorghanBarony{}

func init() {
	f := BBashqorghan巴什库尔干.(*巴什库尔干BashqorghanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bashqorghan",
		TitleName: "巴什库尔干",
		TitleCode: "b_bashqorghan",
	}
}
