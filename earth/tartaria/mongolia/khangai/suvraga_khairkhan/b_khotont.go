package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫吞特KhotontBarony struct {
	feud.BaseBarony
}

var BKhotont赫吞特 feud.Barony = &赫吞特KhotontBarony{}

func init() {
    f := BKhotont赫吞特.(*赫吞特KhotontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khotont",
		TitleName: "赫吞特",
		TitleCode: "b_khotont",
	}
}
