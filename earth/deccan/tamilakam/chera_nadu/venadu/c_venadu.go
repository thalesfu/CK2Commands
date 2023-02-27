package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VenaduCounty interface {
    feud.County
    BAnantasayanam阿难多娑耶南摩() 	feud.Barony
    BAsaudah阿骚陀() 	feud.Barony
    BKandalur甘陀卢尔() 	feud.Barony
    BKanya_kumari羯若鸠摩利() 	feud.Barony
    BKollam故临() 	feud.Barony
    BNagercoil纳盖科伊尔() 	feud.Barony
    BVizhinjam毗只旬() 	feud.Barony
}

type 吠那菟VenaduCounty struct {
	feud.BaseCounty
	阿难多娑耶南摩Anantasayanam 	feud.Barony
	阿骚陀Asaudah 	feud.Barony
	甘陀卢尔Kandalur 	feud.Barony
	羯若鸠摩利Kanya_kumari 	feud.Barony
	故临Kollam 	feud.Barony
	纳盖科伊尔Nagercoil 	feud.Barony
	毗只旬Vizhinjam 	feud.Barony
}

func (c *吠那菟VenaduCounty) BAnantasayanam阿难多娑耶南摩() feud.Barony {
	return c.阿难多娑耶南摩Anantasayanam
}
    
func (c *吠那菟VenaduCounty) BAsaudah阿骚陀() feud.Barony {
	return c.阿骚陀Asaudah
}
    
func (c *吠那菟VenaduCounty) BKandalur甘陀卢尔() feud.Barony {
	return c.甘陀卢尔Kandalur
}
    
func (c *吠那菟VenaduCounty) BKanya_kumari羯若鸠摩利() feud.Barony {
	return c.羯若鸠摩利Kanya_kumari
}
    
func (c *吠那菟VenaduCounty) BKollam故临() feud.Barony {
	return c.故临Kollam
}
    
func (c *吠那菟VenaduCounty) BNagercoil纳盖科伊尔() feud.Barony {
	return c.纳盖科伊尔Nagercoil
}
    
func (c *吠那菟VenaduCounty) BVizhinjam毗只旬() feud.Barony {
	return c.毗只旬Vizhinjam
}
    
var CVenadu吠那菟 VenaduCounty = &吠那菟VenaduCounty{}

func init() {
	f := CVenadu吠那菟.(*吠那菟VenaduCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1413",
		Title:     "venadu",
		TitleName: "吠那菟",
		TitleCode: "c_venadu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿难多娑耶南摩Anantasayanam = BAnantasayanam阿难多娑耶南摩
	f.阿难多娑耶南摩Anantasayanam.SetParent(f)
	
	f.阿骚陀Asaudah = BAsaudah阿骚陀
	f.阿骚陀Asaudah.SetParent(f)
	
	f.甘陀卢尔Kandalur = BKandalur甘陀卢尔
	f.甘陀卢尔Kandalur.SetParent(f)
	
	f.羯若鸠摩利Kanya_kumari = BKanya_kumari羯若鸠摩利
	f.羯若鸠摩利Kanya_kumari.SetParent(f)
	
	f.故临Kollam = BKollam故临
	f.故临Kollam.SetParent(f)
	
	f.纳盖科伊尔Nagercoil = BNagercoil纳盖科伊尔
	f.纳盖科伊尔Nagercoil.SetParent(f)
	
	f.毗只旬Vizhinjam = BVizhinjam毗只旬
	f.毗只旬Vizhinjam.SetParent(f)
	
}
