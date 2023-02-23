package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘陀罗KhudalaBarony struct {
	feud.BaseBarony
}

var BKhudala丘陀罗 feud.Barony = &丘陀罗KhudalaBarony{}

func init() {
	f := BKhudala丘陀罗.(*丘陀罗KhudalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khudala",
		TitleName: "丘陀罗",
		TitleCode: "b_khudala",
	}
}
