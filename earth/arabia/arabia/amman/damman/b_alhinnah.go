package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欣纳赫AlhinnahBarony struct {
	feud.BaseBarony
}

var BAlhinnah欣纳赫 feud.Barony = &欣纳赫AlhinnahBarony{}

func init() {
    f := BAlhinnah欣纳赫.(*欣纳赫AlhinnahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhinnah",
		TitleName: "欣纳赫",
		TitleCode: "b_alhinnah",
	}
}
