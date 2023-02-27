package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布凯马勒BukamalBarony struct {
	feud.BaseBarony
}

var BBukamal布凯马勒 feud.Barony = &布凯马勒BukamalBarony{}

func init() {
    f := BBukamal布凯马勒.(*布凯马勒BukamalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bukamal",
		TitleName: "布凯马勒",
		TitleCode: "b_bukamal",
	}
}
