package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦赫河畔比斯特里察PovazskabystricaBarony struct {
	feud.BaseBarony
}

var BPovazskabystrica瓦赫河畔比斯特里察 feud.Barony = &瓦赫河畔比斯特里察PovazskabystricaBarony{}

func init() {
	f := BPovazskabystrica瓦赫河畔比斯特里察.(*瓦赫河畔比斯特里察PovazskabystricaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "povazskabystrica",
		TitleName: "瓦赫河畔比斯特里察",
		TitleCode: "b_povazskabystrica",
	}
}
