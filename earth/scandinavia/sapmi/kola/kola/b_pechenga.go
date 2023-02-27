package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩琴加PechengaBarony struct {
	feud.BaseBarony
}

var BPechenga佩琴加 feud.Barony = &佩琴加PechengaBarony{}

func init() {
    f := BPechenga佩琴加.(*佩琴加PechengaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pechenga",
		TitleName: "佩琴加",
		TitleCode: "b_pechenga",
	}
}
