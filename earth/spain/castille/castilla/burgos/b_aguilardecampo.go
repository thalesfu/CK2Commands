package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿吉拉尔德坎波奥AguilardecampoBarony struct {
	feud.BaseBarony
}

var BAguilardecampo阿吉拉尔德坎波奥 feud.Barony = &阿吉拉尔德坎波奥AguilardecampoBarony{}

func init() {
	f := BAguilardecampo阿吉拉尔德坎波奥.(*阿吉拉尔德坎波奥AguilardecampoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aguilardecampo",
		TitleName: "阿吉拉尔德坎波奥",
		TitleCode: "b_aguilardecampo",
	}
}
