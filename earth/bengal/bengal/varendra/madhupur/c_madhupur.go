package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MadhupurCounty interface {
    feud.County
    BBaghabarihat婆迦婆利诃() 	feud.Barony
    BBogra菩迦罗() 	feud.Barony
    BKursu鸠须() 	feud.Barony
    BKutra鸠多罗() 	feud.Barony
    BLahpatra罗波多罗() 	feud.Barony
    BMadhupur摩度补罗() 	feud.Barony
    BPabna波婆那() 	feud.Barony
}

type 摩度补罗MadhupurCounty struct {
	feud.BaseCounty
	婆迦婆利诃Baghabarihat 	feud.Barony
	菩迦罗Bogra 	feud.Barony
	鸠须Kursu 	feud.Barony
	鸠多罗Kutra 	feud.Barony
	罗波多罗Lahpatra 	feud.Barony
	摩度补罗Madhupur 	feud.Barony
	波婆那Pabna 	feud.Barony
}

func (c *摩度补罗MadhupurCounty) BBaghabarihat婆迦婆利诃() feud.Barony {
	return c.婆迦婆利诃Baghabarihat
}
    
func (c *摩度补罗MadhupurCounty) BBogra菩迦罗() feud.Barony {
	return c.菩迦罗Bogra
}
    
func (c *摩度补罗MadhupurCounty) BKursu鸠须() feud.Barony {
	return c.鸠须Kursu
}
    
func (c *摩度补罗MadhupurCounty) BKutra鸠多罗() feud.Barony {
	return c.鸠多罗Kutra
}
    
func (c *摩度补罗MadhupurCounty) BLahpatra罗波多罗() feud.Barony {
	return c.罗波多罗Lahpatra
}
    
func (c *摩度补罗MadhupurCounty) BMadhupur摩度补罗() feud.Barony {
	return c.摩度补罗Madhupur
}
    
func (c *摩度补罗MadhupurCounty) BPabna波婆那() feud.Barony {
	return c.波婆那Pabna
}
    
var CMadhupur摩度补罗 MadhupurCounty = &摩度补罗MadhupurCounty{}

func init() {
	f := CMadhupur摩度补罗.(*摩度补罗MadhupurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1325",
		Title:     "madhupur",
		TitleName: "摩度补罗",
		TitleCode: "c_madhupur",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆迦婆利诃Baghabarihat = BBaghabarihat婆迦婆利诃
	f.婆迦婆利诃Baghabarihat.SetParent(f)
	
	f.菩迦罗Bogra = BBogra菩迦罗
	f.菩迦罗Bogra.SetParent(f)
	
	f.鸠须Kursu = BKursu鸠须
	f.鸠须Kursu.SetParent(f)
	
	f.鸠多罗Kutra = BKutra鸠多罗
	f.鸠多罗Kutra.SetParent(f)
	
	f.罗波多罗Lahpatra = BLahpatra罗波多罗
	f.罗波多罗Lahpatra.SetParent(f)
	
	f.摩度补罗Madhupur = BMadhupur摩度补罗
	f.摩度补罗Madhupur.SetParent(f)
	
	f.波婆那Pabna = BPabna波婆那
	f.波婆那Pabna.SetParent(f)
	
}
