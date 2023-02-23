package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾汉帕那赫JahanpanahBarony struct {
	feud.BaseBarony
}

var BJahanpanah贾汉帕那赫 feud.Barony = &贾汉帕那赫JahanpanahBarony{}

func init() {
	f := BJahanpanah贾汉帕那赫.(*贾汉帕那赫JahanpanahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jahanpanah",
		TitleName: "贾汉帕那赫",
		TitleCode: "b_jahanpanah",
	}
}
