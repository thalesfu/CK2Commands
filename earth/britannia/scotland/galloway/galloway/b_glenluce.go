package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格伦卢斯GlenluceBarony struct {
	feud.BaseBarony
}

var BGlenluce格伦卢斯 feud.Barony = &格伦卢斯GlenluceBarony{}

func init() {
	f := BGlenluce格伦卢斯.(*格伦卢斯GlenluceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glenluce",
		TitleName: "格伦卢斯",
		TitleCode: "b_glenluce",
	}
}
