package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提斯阿姆提TisiamtiBarony struct {
	feud.BaseBarony
}

var BTisiamti提斯阿姆提 feud.Barony = &提斯阿姆提TisiamtiBarony{}

func init() {
	f := BTisiamti提斯阿姆提.(*提斯阿姆提TisiamtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tisiamti",
		TitleName: "提斯阿姆提",
		TitleCode: "b_tisiamti",
	}
}
