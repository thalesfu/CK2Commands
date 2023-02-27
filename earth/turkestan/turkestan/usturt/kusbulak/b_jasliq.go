package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎斯雷克JasliqBarony struct {
	feud.BaseBarony
}

var BJasliq扎斯雷克 feud.Barony = &扎斯雷克JasliqBarony{}

func init() {
    f := BJasliq扎斯雷克.(*扎斯雷克JasliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jasliq",
		TitleName: "扎斯雷克",
		TitleCode: "b_jasliq",
	}
}
