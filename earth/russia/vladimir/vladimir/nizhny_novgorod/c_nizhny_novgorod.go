package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Nizhny_novgorodCounty interface {
    feud.County
    BBalakhna巴拉赫纳() 	feud.Barony
    BBogorodsk博戈罗茨克() 	feud.Barony
    BKnyaginino克尼亚吉尼诺() 	feud.Barony
    BKstovo克斯托沃() 	feud.Barony
    BLyskovo雷斯科沃() 	feud.Barony
    BNizhnynovgorod下诺夫哥罗德() 	feud.Barony
    BPavlovo巴甫洛沃() 	feud.Barony
}

type 下诺夫哥罗德Nizhny_novgorodCounty struct {
	feud.BaseCounty
	巴拉赫纳Balakhna 	feud.Barony
	博戈罗茨克Bogorodsk 	feud.Barony
	克尼亚吉尼诺Knyaginino 	feud.Barony
	克斯托沃Kstovo 	feud.Barony
	雷斯科沃Lyskovo 	feud.Barony
	下诺夫哥罗德Nizhnynovgorod 	feud.Barony
	巴甫洛沃Pavlovo 	feud.Barony
}

func (c *下诺夫哥罗德Nizhny_novgorodCounty) BBalakhna巴拉赫纳() feud.Barony {
	return c.巴拉赫纳Balakhna
}
    
func (c *下诺夫哥罗德Nizhny_novgorodCounty) BBogorodsk博戈罗茨克() feud.Barony {
	return c.博戈罗茨克Bogorodsk
}
    
func (c *下诺夫哥罗德Nizhny_novgorodCounty) BKnyaginino克尼亚吉尼诺() feud.Barony {
	return c.克尼亚吉尼诺Knyaginino
}
    
func (c *下诺夫哥罗德Nizhny_novgorodCounty) BKstovo克斯托沃() feud.Barony {
	return c.克斯托沃Kstovo
}
    
func (c *下诺夫哥罗德Nizhny_novgorodCounty) BLyskovo雷斯科沃() feud.Barony {
	return c.雷斯科沃Lyskovo
}
    
func (c *下诺夫哥罗德Nizhny_novgorodCounty) BNizhnynovgorod下诺夫哥罗德() feud.Barony {
	return c.下诺夫哥罗德Nizhnynovgorod
}
    
func (c *下诺夫哥罗德Nizhny_novgorodCounty) BPavlovo巴甫洛沃() feud.Barony {
	return c.巴甫洛沃Pavlovo
}
    
var CNizhny_novgorod下诺夫哥罗德 Nizhny_novgorodCounty = &下诺夫哥罗德Nizhny_novgorodCounty{}

func init() {
	f := CNizhny_novgorod下诺夫哥罗德.(*下诺夫哥罗德Nizhny_novgorodCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "584",
		Title:     "nizhny_novgorod",
		TitleName: "下诺夫哥罗德",
		TitleCode: "c_nizhny_novgorod",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉赫纳Balakhna = BBalakhna巴拉赫纳
	f.巴拉赫纳Balakhna.SetParent(f)
	
	f.博戈罗茨克Bogorodsk = BBogorodsk博戈罗茨克
	f.博戈罗茨克Bogorodsk.SetParent(f)
	
	f.克尼亚吉尼诺Knyaginino = BKnyaginino克尼亚吉尼诺
	f.克尼亚吉尼诺Knyaginino.SetParent(f)
	
	f.克斯托沃Kstovo = BKstovo克斯托沃
	f.克斯托沃Kstovo.SetParent(f)
	
	f.雷斯科沃Lyskovo = BLyskovo雷斯科沃
	f.雷斯科沃Lyskovo.SetParent(f)
	
	f.下诺夫哥罗德Nizhnynovgorod = BNizhnynovgorod下诺夫哥罗德
	f.下诺夫哥罗德Nizhnynovgorod.SetParent(f)
	
	f.巴甫洛沃Pavlovo = BPavlovo巴甫洛沃
	f.巴甫洛沃Pavlovo.SetParent(f)
	
}
