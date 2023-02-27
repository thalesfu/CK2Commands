package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉祥萨迦PelsakyaBarony struct {
	feud.BaseBarony
}

var BPelsakya吉祥萨迦 feud.Barony = &吉祥萨迦PelsakyaBarony{}

func init() {
    f := BPelsakya吉祥萨迦.(*吉祥萨迦PelsakyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pelsakya",
		TitleName: "吉祥萨迦",
		TitleCode: "b_pelsakya",
	}
}
