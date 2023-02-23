package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆塞德LemseiedBarony struct {
	feud.BaseBarony
}

var BLemseied姆塞德 feud.Barony = &姆塞德LemseiedBarony{}

func init() {
	f := BLemseied姆塞德.(*姆塞德LemseiedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lemseied",
		TitleName: "姆塞德",
		TitleCode: "b_lemseied",
	}
}
