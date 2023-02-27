package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 强贝JambayBarony struct {
	feud.BaseBarony
}

var BJambay强贝 feud.Barony = &强贝JambayBarony{}

func init() {
    f := BJambay强贝.(*强贝JambayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jambay",
		TitleName: "强贝",
		TitleCode: "b_jambay",
	}
}
