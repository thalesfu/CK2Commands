package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆纳桑Umm_al_nasanBarony struct {
	feud.BaseBarony
}

var BUmm_al_nasan乌姆纳桑 feud.Barony = &乌姆纳桑Umm_al_nasanBarony{}

func init() {
    f := BUmm_al_nasan乌姆纳桑.(*乌姆纳桑Umm_al_nasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umm_al_nasan",
		TitleName: "乌姆纳桑",
		TitleCode: "b_umm_al_nasan",
	}
}
