package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎敏达瓦尔ZamindawarBarony struct {
	feud.BaseBarony
}

var BZamindawar扎敏达瓦尔 feud.Barony = &扎敏达瓦尔ZamindawarBarony{}

func init() {
	f := BZamindawar扎敏达瓦尔.(*扎敏达瓦尔ZamindawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zamindawar",
		TitleName: "扎敏达瓦尔",
		TitleCode: "b_zamindawar",
	}
}
