package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马伊MaillyBarony struct {
	feud.BaseBarony
}

var BMailly马伊 feud.Barony = &马伊MaillyBarony{}

func init() {
    f := BMailly马伊.(*马伊MaillyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mailly",
		TitleName: "马伊",
		TitleCode: "b_mailly",
	}
}
