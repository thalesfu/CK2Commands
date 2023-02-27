package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉木乃JeminayBarony struct {
	feud.BaseBarony
}

var BJeminay吉木乃 feud.Barony = &吉木乃JeminayBarony{}

func init() {
    f := BJeminay吉木乃.(*吉木乃JeminayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jeminay",
		TitleName: "吉木乃",
		TitleCode: "b_jeminay",
	}
}
