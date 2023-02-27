package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SennarCounty interface {
    feud.County
    BBakkah巴卡赫() 	feud.Barony
    BGalgani格甘尼() 	feud.Barony
    BKinanah吉娜赫() 	feud.Barony
    BKukur库库尔() 	feud.Barony
    BSennar森纳尔() 	feud.Barony
    BSinjah辛贾() 	feud.Barony
    BTabat塔巴特() 	feud.Barony
    BWadrawah瓦德拉瓦赫() 	feud.Barony
}

type 森纳尔SennarCounty struct {
	feud.BaseCounty
	巴卡赫Bakkah 	feud.Barony
	格甘尼Galgani 	feud.Barony
	吉娜赫Kinanah 	feud.Barony
	库库尔Kukur 	feud.Barony
	森纳尔Sennar 	feud.Barony
	辛贾Sinjah 	feud.Barony
	塔巴特Tabat 	feud.Barony
	瓦德拉瓦赫Wadrawah 	feud.Barony
}

func (c *森纳尔SennarCounty) BBakkah巴卡赫() feud.Barony {
	return c.巴卡赫Bakkah
}
    
func (c *森纳尔SennarCounty) BGalgani格甘尼() feud.Barony {
	return c.格甘尼Galgani
}
    
func (c *森纳尔SennarCounty) BKinanah吉娜赫() feud.Barony {
	return c.吉娜赫Kinanah
}
    
func (c *森纳尔SennarCounty) BKukur库库尔() feud.Barony {
	return c.库库尔Kukur
}
    
func (c *森纳尔SennarCounty) BSennar森纳尔() feud.Barony {
	return c.森纳尔Sennar
}
    
func (c *森纳尔SennarCounty) BSinjah辛贾() feud.Barony {
	return c.辛贾Sinjah
}
    
func (c *森纳尔SennarCounty) BTabat塔巴特() feud.Barony {
	return c.塔巴特Tabat
}
    
func (c *森纳尔SennarCounty) BWadrawah瓦德拉瓦赫() feud.Barony {
	return c.瓦德拉瓦赫Wadrawah
}
    
var CSennar森纳尔 SennarCounty = &森纳尔SennarCounty{}

func init() {
	f := CSennar森纳尔.(*森纳尔SennarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "880",
		Title:     "sennar",
		TitleName: "森纳尔",
		TitleCode: "c_sennar",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴卡赫Bakkah = BBakkah巴卡赫
	f.巴卡赫Bakkah.SetParent(f)
	
	f.格甘尼Galgani = BGalgani格甘尼
	f.格甘尼Galgani.SetParent(f)
	
	f.吉娜赫Kinanah = BKinanah吉娜赫
	f.吉娜赫Kinanah.SetParent(f)
	
	f.库库尔Kukur = BKukur库库尔
	f.库库尔Kukur.SetParent(f)
	
	f.森纳尔Sennar = BSennar森纳尔
	f.森纳尔Sennar.SetParent(f)
	
	f.辛贾Sinjah = BSinjah辛贾
	f.辛贾Sinjah.SetParent(f)
	
	f.塔巴特Tabat = BTabat塔巴特
	f.塔巴特Tabat.SetParent(f)
	
	f.瓦德拉瓦赫Wadrawah = BWadrawah瓦德拉瓦赫
	f.瓦德拉瓦赫Wadrawah.SetParent(f)
	
}
