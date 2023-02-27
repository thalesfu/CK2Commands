package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑞塞JusseyBarony struct {
	feud.BaseBarony
}

var BJussey瑞塞 feud.Barony = &瑞塞JusseyBarony{}

func init() {
    f := BJussey瑞塞.(*瑞塞JusseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jussey",
		TitleName: "瑞塞",
		TitleCode: "b_jussey",
	}
}
