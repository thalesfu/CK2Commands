package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎拜萨马拉JabalsamraBarony struct {
	feud.BaseBarony
}

var BJabalsamra扎拜萨马拉 feud.Barony = &扎拜萨马拉JabalsamraBarony{}

func init() {
    f := BJabalsamra扎拜萨马拉.(*扎拜萨马拉JabalsamraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jabalsamra",
		TitleName: "扎拜萨马拉",
		TitleCode: "b_jabalsamra",
	}
}
