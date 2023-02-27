package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文格鲁夫WegrowBarony struct {
	feud.BaseBarony
}

var BWegrow文格鲁夫 feud.Barony = &文格鲁夫WegrowBarony{}

func init() {
    f := BWegrow文格鲁夫.(*文格鲁夫WegrowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wegrow",
		TitleName: "文格鲁夫",
		TitleCode: "b_wegrow",
	}
}
