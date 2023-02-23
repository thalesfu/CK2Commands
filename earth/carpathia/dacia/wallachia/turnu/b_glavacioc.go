package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉瓦乔克GlavaciocBarony struct {
	feud.BaseBarony
}

var BGlavacioc格拉瓦乔克 feud.Barony = &格拉瓦乔克GlavaciocBarony{}

func init() {
	f := BGlavacioc格拉瓦乔克.(*格拉瓦乔克GlavaciocBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glavacioc",
		TitleName: "格拉瓦乔克",
		TitleCode: "b_glavacioc",
	}
}
