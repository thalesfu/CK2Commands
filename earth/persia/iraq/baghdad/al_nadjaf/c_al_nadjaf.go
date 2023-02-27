package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_nadjafCounty interface {
    feud.County
    BJaarah扎阿尔赫() 	feud.Barony
    BJasim贾西姆() 	feud.Barony
    BKufah库费() 	feud.Barony
    BMidhrawi米德霍拉维() 	feud.Barony
    BNadjaf纳杰夫() 	feud.Barony
    BRashid罗塞塔() 	feud.Barony
    BTaqtaqanah塔达喀纳赫() 	feud.Barony
}

type 纳杰夫Al_nadjafCounty struct {
	feud.BaseCounty
	扎阿尔赫Jaarah 	feud.Barony
	贾西姆Jasim 	feud.Barony
	库费Kufah 	feud.Barony
	米德霍拉维Midhrawi 	feud.Barony
	纳杰夫Nadjaf 	feud.Barony
	罗塞塔Rashid 	feud.Barony
	塔达喀纳赫Taqtaqanah 	feud.Barony
}

func (c *纳杰夫Al_nadjafCounty) BJaarah扎阿尔赫() feud.Barony {
	return c.扎阿尔赫Jaarah
}
    
func (c *纳杰夫Al_nadjafCounty) BJasim贾西姆() feud.Barony {
	return c.贾西姆Jasim
}
    
func (c *纳杰夫Al_nadjafCounty) BKufah库费() feud.Barony {
	return c.库费Kufah
}
    
func (c *纳杰夫Al_nadjafCounty) BMidhrawi米德霍拉维() feud.Barony {
	return c.米德霍拉维Midhrawi
}
    
func (c *纳杰夫Al_nadjafCounty) BNadjaf纳杰夫() feud.Barony {
	return c.纳杰夫Nadjaf
}
    
func (c *纳杰夫Al_nadjafCounty) BRashid罗塞塔() feud.Barony {
	return c.罗塞塔Rashid
}
    
func (c *纳杰夫Al_nadjafCounty) BTaqtaqanah塔达喀纳赫() feud.Barony {
	return c.塔达喀纳赫Taqtaqanah
}
    
var CAl_nadjaf纳杰夫 Al_nadjafCounty = &纳杰夫Al_nadjafCounty{}

func init() {
	f := CAl_nadjaf纳杰夫.(*纳杰夫Al_nadjafCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "692",
		Title:     "al_nadjaf",
		TitleName: "纳杰夫",
		TitleCode: "c_al_nadjaf",
		Baronies:  map[string]feud.Barony{},
	}

	f.扎阿尔赫Jaarah = BJaarah扎阿尔赫
	f.扎阿尔赫Jaarah.SetParent(f)
	
	f.贾西姆Jasim = BJasim贾西姆
	f.贾西姆Jasim.SetParent(f)
	
	f.库费Kufah = BKufah库费
	f.库费Kufah.SetParent(f)
	
	f.米德霍拉维Midhrawi = BMidhrawi米德霍拉维
	f.米德霍拉维Midhrawi.SetParent(f)
	
	f.纳杰夫Nadjaf = BNadjaf纳杰夫
	f.纳杰夫Nadjaf.SetParent(f)
	
	f.罗塞塔Rashid = BRashid罗塞塔
	f.罗塞塔Rashid.SetParent(f)
	
	f.塔达喀纳赫Taqtaqanah = BTaqtaqanah塔达喀纳赫
	f.塔达喀纳赫Taqtaqanah.SetParent(f)
	
}
