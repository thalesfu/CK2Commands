package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TuggurtCounty interface {
    feud.County
    BElalia阿利耶() 	feud.Barony
    BElbour布尔() 	feud.Barony
    BMeftah梅夫塔() 	feud.Barony
    BSmadja斯马贾() 	feud.Barony
    BTennsaout滕绍特() 	feud.Barony
    BTimeliline蒂梅利林() 	feud.Barony
    BTuggurt图古尔特() 	feud.Barony
}

type 图古尔特TuggurtCounty struct {
	feud.BaseCounty
	阿利耶Elalia 	feud.Barony
	布尔Elbour 	feud.Barony
	梅夫塔Meftah 	feud.Barony
	斯马贾Smadja 	feud.Barony
	滕绍特Tennsaout 	feud.Barony
	蒂梅利林Timeliline 	feud.Barony
	图古尔特Tuggurt 	feud.Barony
}

func (c *图古尔特TuggurtCounty) BElalia阿利耶() feud.Barony {
	return c.阿利耶Elalia
}
    
func (c *图古尔特TuggurtCounty) BElbour布尔() feud.Barony {
	return c.布尔Elbour
}
    
func (c *图古尔特TuggurtCounty) BMeftah梅夫塔() feud.Barony {
	return c.梅夫塔Meftah
}
    
func (c *图古尔特TuggurtCounty) BSmadja斯马贾() feud.Barony {
	return c.斯马贾Smadja
}
    
func (c *图古尔特TuggurtCounty) BTennsaout滕绍特() feud.Barony {
	return c.滕绍特Tennsaout
}
    
func (c *图古尔特TuggurtCounty) BTimeliline蒂梅利林() feud.Barony {
	return c.蒂梅利林Timeliline
}
    
func (c *图古尔特TuggurtCounty) BTuggurt图古尔特() feud.Barony {
	return c.图古尔特Tuggurt
}
    
var CTuggurt图古尔特 TuggurtCounty = &图古尔特TuggurtCounty{}

func init() {
	f := CTuggurt图古尔特.(*图古尔特TuggurtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1729",
		Title:     "tuggurt",
		TitleName: "图古尔特",
		TitleCode: "c_tuggurt",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿利耶Elalia = BElalia阿利耶
	f.阿利耶Elalia.SetParent(f)
	
	f.布尔Elbour = BElbour布尔
	f.布尔Elbour.SetParent(f)
	
	f.梅夫塔Meftah = BMeftah梅夫塔
	f.梅夫塔Meftah.SetParent(f)
	
	f.斯马贾Smadja = BSmadja斯马贾
	f.斯马贾Smadja.SetParent(f)
	
	f.滕绍特Tennsaout = BTennsaout滕绍特
	f.滕绍特Tennsaout.SetParent(f)
	
	f.蒂梅利林Timeliline = BTimeliline蒂梅利林
	f.蒂梅利林Timeliline.SetParent(f)
	
	f.图古尔特Tuggurt = BTuggurt图古尔特
	f.图古尔特Tuggurt.SetParent(f)
	
}
