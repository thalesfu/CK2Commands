package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 狄奥多西波利斯TheodosiopolisBarony struct {
	feud.BaseBarony
}

var BTheodosiopolis狄奥多西波利斯 feud.Barony = &狄奥多西波利斯TheodosiopolisBarony{}

func init() {
	f := BTheodosiopolis狄奥多西波利斯.(*狄奥多西波利斯TheodosiopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "theodosiopolis",
		TitleName: "狄奥多西波利斯",
		TitleCode: "b_theodosiopolis",
	}
}
