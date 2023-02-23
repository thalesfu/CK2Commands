package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库西里KussiriBarony struct {
	feud.BaseBarony
}

var BKussiri库西里 feud.Barony = &库西里KussiriBarony{}

func init() {
	f := BKussiri库西里.(*库西里KussiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kussiri",
		TitleName: "库西里",
		TitleCode: "b_kussiri",
	}
}
