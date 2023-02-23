package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑诃斯罗婆呼SahastrabahuBarony struct {
	feud.BaseBarony
}

var BSahastrabahu娑诃斯罗婆呼 feud.Barony = &娑诃斯罗婆呼SahastrabahuBarony{}

func init() {
	f := BSahastrabahu娑诃斯罗婆呼.(*娑诃斯罗婆呼SahastrabahuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sahastrabahu",
		TitleName: "娑诃斯罗婆呼",
		TitleCode: "b_sahastrabahu",
	}
}
