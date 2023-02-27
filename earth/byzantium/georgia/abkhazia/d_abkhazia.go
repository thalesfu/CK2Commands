package abkhazia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/abkhazia/abkhazia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/abkhazia/egrisi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbkhaziaDuke interface {
    feud.Duke
    CAbkhazia阿布哈兹() 	abkhazia.AbkhaziaCounty
    CEgrisi埃格里西() 	egrisi.EgrisiCounty
}

type 阿布哈兹AbkhaziaDuke struct {
	feud.BaseDuke
	阿布哈兹Abkhazia 	abkhazia.AbkhaziaCounty
	埃格里西Egrisi 	egrisi.EgrisiCounty
}

func (d *阿布哈兹AbkhaziaDuke) CAbkhazia阿布哈兹() abkhazia.AbkhaziaCounty {
	return d.阿布哈兹Abkhazia
}
    
func (d *阿布哈兹AbkhaziaDuke) CEgrisi埃格里西() egrisi.EgrisiCounty {
	return d.埃格里西Egrisi
}
    
var DAbkhazia阿布哈兹 AbkhaziaDuke = &阿布哈兹AbkhaziaDuke{}

func init() {
	f := DAbkhazia阿布哈兹.(*阿布哈兹AbkhaziaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "abkhazia",
		TitleName: "阿布哈兹",
		TitleCode: "d_abkhazia",
		Counties:  map[string]feud.County{},
	}

	f.阿布哈兹Abkhazia = abkhazia.CAbkhazia阿布哈兹
	f.阿布哈兹Abkhazia.SetParent(f)
	
	f.埃格里西Egrisi = egrisi.CEgrisi埃格里西
	f.埃格里西Egrisi.SetParent(f)
	
}
