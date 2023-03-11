package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅罗斯拉夫JaroslavBarony struct {
	feud.BaseBarony
}

var BJaroslav雅罗斯拉夫 feud.Barony = &雅罗斯拉夫JaroslavBarony{}

func init() {
    f := BJaroslav雅罗斯拉夫.(*雅罗斯拉夫JaroslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaroslav",
		TitleName: "雅罗斯拉夫",
		TitleCode: "b_jaroslav",
	}
}
