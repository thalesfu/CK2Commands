package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里阿佐夫西克PryazovskeBarony struct {
	feud.BaseBarony
}

var BPryazovske普里阿佐夫西克 feud.Barony = &普里阿佐夫西克PryazovskeBarony{}

func init() {
    f := BPryazovske普里阿佐夫西克.(*普里阿佐夫西克PryazovskeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pryazovske",
		TitleName: "普里阿佐夫西克",
		TitleCode: "b_pryazovske",
	}
}
