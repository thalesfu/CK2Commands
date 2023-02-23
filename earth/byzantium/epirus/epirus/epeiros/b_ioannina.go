package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约阿尼纳IoanninaBarony struct {
	feud.BaseBarony
}

var BIoannina约阿尼纳 feud.Barony = &约阿尼纳IoanninaBarony{}

func init() {
	f := BIoannina约阿尼纳.(*约阿尼纳IoanninaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ioannina",
		TitleName: "约阿尼纳",
		TitleCode: "b_ioannina",
	}
}
