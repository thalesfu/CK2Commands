package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米库林MykulynBarony struct {
	feud.BaseBarony
}

var BMykulyn米库林 feud.Barony = &米库林MykulynBarony{}

func init() {
    f := BMykulyn米库林.(*米库林MykulynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mykulyn",
		TitleName: "米库林",
		TitleCode: "b_mykulyn",
	}
}
