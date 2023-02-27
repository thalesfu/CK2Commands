package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺沃亚佐夫斯克NovoazovskBarony struct {
	feud.BaseBarony
}

var BNovoazovsk诺沃亚佐夫斯克 feud.Barony = &诺沃亚佐夫斯克NovoazovskBarony{}

func init() {
    f := BNovoazovsk诺沃亚佐夫斯克.(*诺沃亚佐夫斯克NovoazovskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novoazovsk",
		TitleName: "诺沃亚佐夫斯克",
		TitleCode: "b_novoazovsk",
	}
}
