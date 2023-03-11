package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢茨克LutskBarony struct {
	feud.BaseBarony
}

var BLutsk卢茨克 feud.Barony = &卢茨克LutskBarony{}

func init() {
    f := BLutsk卢茨克.(*卢茨克LutskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lutsk",
		TitleName: "卢茨克",
		TitleCode: "b_lutsk",
	}
}
