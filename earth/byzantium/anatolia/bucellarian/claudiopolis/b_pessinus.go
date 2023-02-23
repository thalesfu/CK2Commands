package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 培希努PessinusBarony struct {
	feud.BaseBarony
}

var BPessinus培希努 feud.Barony = &培希努PessinusBarony{}

func init() {
	f := BPessinus培希努.(*培希努PessinusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pessinus",
		TitleName: "培希努",
		TitleCode: "b_pessinus",
	}
}
