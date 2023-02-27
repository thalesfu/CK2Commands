package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷文GrevenBarony struct {
	feud.BaseBarony
}

var BGreven格雷文 feud.Barony = &格雷文GrevenBarony{}

func init() {
    f := BGreven格雷文.(*格雷文GrevenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "greven",
		TitleName: "格雷文",
		TitleCode: "b_greven",
	}
}
