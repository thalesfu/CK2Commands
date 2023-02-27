package pommerania

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pommerania/neumark"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pommerania/stettin"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/pommerania/wolgast"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PommeraniaDuke interface {
    feud.Duke
    CNeumark诺伊马尔克() 	neumark.NeumarkCounty
    CStettin斯德丁() 	stettin.StettinCounty
    CWolgast沃尔加斯特() 	wolgast.WolgastCounty
}

type 波美拉尼亚PommeraniaDuke struct {
	feud.BaseDuke
	诺伊马尔克Neumark 	neumark.NeumarkCounty
	斯德丁Stettin 	stettin.StettinCounty
	沃尔加斯特Wolgast 	wolgast.WolgastCounty
}

func (d *波美拉尼亚PommeraniaDuke) CNeumark诺伊马尔克() neumark.NeumarkCounty {
	return d.诺伊马尔克Neumark
}
    
func (d *波美拉尼亚PommeraniaDuke) CStettin斯德丁() stettin.StettinCounty {
	return d.斯德丁Stettin
}
    
func (d *波美拉尼亚PommeraniaDuke) CWolgast沃尔加斯特() wolgast.WolgastCounty {
	return d.沃尔加斯特Wolgast
}
    
var DPommerania波美拉尼亚 PommeraniaDuke = &波美拉尼亚PommeraniaDuke{}

func init() {
	f := DPommerania波美拉尼亚.(*波美拉尼亚PommeraniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pommerania",
		TitleName: "波美拉尼亚",
		TitleCode: "d_pommerania",
		Counties:  map[string]feud.County{},
	}

	f.诺伊马尔克Neumark = neumark.CNeumark诺伊马尔克
	f.诺伊马尔克Neumark.SetParent(f)
	
	f.斯德丁Stettin = stettin.CStettin斯德丁
	f.斯德丁Stettin.SetParent(f)
	
	f.沃尔加斯特Wolgast = wolgast.CWolgast沃尔加斯特
	f.沃尔加斯特Wolgast.SetParent(f)
	
}
