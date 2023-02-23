package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格鲁德HrudBarony struct {
	feud.BaseBarony
}

var BHrud格鲁德 feud.Barony = &格鲁德HrudBarony{}

func init() {
	f := BHrud格鲁德.(*格鲁德HrudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hrud",
		TitleName: "格鲁德",
		TitleCode: "b_hrud",
	}
}
