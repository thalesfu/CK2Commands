package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆卢赫AmlwchBarony struct {
	feud.BaseBarony
}

var BAmlwch阿姆卢赫 feud.Barony = &阿姆卢赫AmlwchBarony{}

func init() {
	f := BAmlwch阿姆卢赫.(*阿姆卢赫AmlwchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amlwch",
		TitleName: "阿姆卢赫",
		TitleCode: "b_amlwch",
	}
}
