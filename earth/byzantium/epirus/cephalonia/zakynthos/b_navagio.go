package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳瓦吉奥NavagioBarony struct {
	feud.BaseBarony
}

var BNavagio纳瓦吉奥 feud.Barony = &纳瓦吉奥NavagioBarony{}

func init() {
    f := BNavagio纳瓦吉奥.(*纳瓦吉奥NavagioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "navagio",
		TitleName: "纳瓦吉奥",
		TitleCode: "b_navagio",
	}
}
