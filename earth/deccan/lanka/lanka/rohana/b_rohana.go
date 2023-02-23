package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢诃拿RohanaBarony struct {
	feud.BaseBarony
}

var BRohana卢诃拿 feud.Barony = &卢诃拿RohanaBarony{}

func init() {
	f := BRohana卢诃拿.(*卢诃拿RohanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohana",
		TitleName: "卢诃拿",
		TitleCode: "b_rohana",
	}
}
