package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆塞班UmmalsabaanBarony struct {
	feud.BaseBarony
}

var BUmmalsabaan乌姆塞班 feud.Barony = &乌姆塞班UmmalsabaanBarony{}

func init() {
	f := BUmmalsabaan乌姆塞班.(*乌姆塞班UmmalsabaanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ummalsabaan",
		TitleName: "乌姆塞班",
		TitleCode: "b_ummalsabaan",
	}
}
