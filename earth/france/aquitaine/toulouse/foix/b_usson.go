package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于松UssonBarony struct {
	feud.BaseBarony
}

var BUsson于松 feud.Barony = &于松UssonBarony{}

func init() {
    f := BUsson于松.(*于松UssonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usson",
		TitleName: "于松",
		TitleCode: "b_usson",
	}
}
