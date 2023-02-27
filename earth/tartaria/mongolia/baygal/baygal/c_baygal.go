package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaygalCounty interface {
    feud.County
    BBargujin八儿忽真() 	feud.Barony
    BBaygal贝加尔() 	feud.Barony
    BDarkhan达尔汗() 	feud.Barony
    BGaluuta嘎鲁特() 	feud.Barony
    BKyakhta恰克图() 	feud.Barony
    BMerkit蔑儿乞() 	feud.Barony
    BUlyun乌柳恩() 	feud.Barony
}

type 贝加尔BaygalCounty struct {
	feud.BaseCounty
	八儿忽真Bargujin 	feud.Barony
	贝加尔Baygal 	feud.Barony
	达尔汗Darkhan 	feud.Barony
	嘎鲁特Galuuta 	feud.Barony
	恰克图Kyakhta 	feud.Barony
	蔑儿乞Merkit 	feud.Barony
	乌柳恩Ulyun 	feud.Barony
}

func (c *贝加尔BaygalCounty) BBargujin八儿忽真() feud.Barony {
	return c.八儿忽真Bargujin
}
    
func (c *贝加尔BaygalCounty) BBaygal贝加尔() feud.Barony {
	return c.贝加尔Baygal
}
    
func (c *贝加尔BaygalCounty) BDarkhan达尔汗() feud.Barony {
	return c.达尔汗Darkhan
}
    
func (c *贝加尔BaygalCounty) BGaluuta嘎鲁特() feud.Barony {
	return c.嘎鲁特Galuuta
}
    
func (c *贝加尔BaygalCounty) BKyakhta恰克图() feud.Barony {
	return c.恰克图Kyakhta
}
    
func (c *贝加尔BaygalCounty) BMerkit蔑儿乞() feud.Barony {
	return c.蔑儿乞Merkit
}
    
func (c *贝加尔BaygalCounty) BUlyun乌柳恩() feud.Barony {
	return c.乌柳恩Ulyun
}
    
var CBaygal贝加尔 BaygalCounty = &贝加尔BaygalCounty{}

func init() {
	f := CBaygal贝加尔.(*贝加尔BaygalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1460",
		Title:     "baygal",
		TitleName: "贝加尔",
		TitleCode: "c_baygal",
		Baronies:  map[string]feud.Barony{},
	}

	f.八儿忽真Bargujin = BBargujin八儿忽真
	f.八儿忽真Bargujin.SetParent(f)
	
	f.贝加尔Baygal = BBaygal贝加尔
	f.贝加尔Baygal.SetParent(f)
	
	f.达尔汗Darkhan = BDarkhan达尔汗
	f.达尔汗Darkhan.SetParent(f)
	
	f.嘎鲁特Galuuta = BGaluuta嘎鲁特
	f.嘎鲁特Galuuta.SetParent(f)
	
	f.恰克图Kyakhta = BKyakhta恰克图
	f.恰克图Kyakhta.SetParent(f)
	
	f.蔑儿乞Merkit = BMerkit蔑儿乞
	f.蔑儿乞Merkit.SetParent(f)
	
	f.乌柳恩Ulyun = BUlyun乌柳恩
	f.乌柳恩Ulyun.SetParent(f)
	
}
