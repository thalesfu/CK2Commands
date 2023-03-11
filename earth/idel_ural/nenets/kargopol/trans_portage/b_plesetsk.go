package trans_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普列谢茨克PlesetskBarony struct {
	feud.BaseBarony
}

var BPlesetsk普列谢茨克 feud.Barony = &普列谢茨克PlesetskBarony{}

func init() {
    f := BPlesetsk普列谢茨克.(*普列谢茨克PlesetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plesetsk",
		TitleName: "普列谢茨克",
		TitleCode: "b_plesetsk",
	}
}
