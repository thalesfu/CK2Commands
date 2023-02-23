package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿库季哈AkutikhaBarony struct {
	feud.BaseBarony
}

var BAkutikha阿库季哈 feud.Barony = &阿库季哈AkutikhaBarony{}

func init() {
	f := BAkutikha阿库季哈.(*阿库季哈AkutikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akutikha",
		TitleName: "阿库季哈",
		TitleCode: "b_akutikha",
	}
}
