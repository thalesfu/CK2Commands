package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呾剌健TalaqanBarony struct {
	feud.BaseBarony
}

var BTalaqan呾剌健 feud.Barony = &呾剌健TalaqanBarony{}

func init() {
    f := BTalaqan呾剌健.(*呾剌健TalaqanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talaqan",
		TitleName: "呾剌健",
		TitleCode: "b_talaqan",
	}
}
