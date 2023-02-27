package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Lower_silesiaCounty interface {
    feud.County
    BBytomodrz奥得河畔比托姆() 	feud.Barony
    BGlogow格沃古夫() 	feud.Barony
    BKrosno克罗斯诺() 	feud.Barony
    BNowogrod诺沃格鲁德() 	feud.Barony
    BSwiebozin希维博津() 	feud.Barony
    BZagan扎甘() 	feud.Barony
    BZielonagora绿山城() 	feud.Barony
}

type 下西里西亚Lower_silesiaCounty struct {
	feud.BaseCounty
	奥得河畔比托姆Bytomodrz 	feud.Barony
	格沃古夫Glogow 	feud.Barony
	克罗斯诺Krosno 	feud.Barony
	诺沃格鲁德Nowogrod 	feud.Barony
	希维博津Swiebozin 	feud.Barony
	扎甘Zagan 	feud.Barony
	绿山城Zielonagora 	feud.Barony
}

func (c *下西里西亚Lower_silesiaCounty) BBytomodrz奥得河畔比托姆() feud.Barony {
	return c.奥得河畔比托姆Bytomodrz
}
    
func (c *下西里西亚Lower_silesiaCounty) BGlogow格沃古夫() feud.Barony {
	return c.格沃古夫Glogow
}
    
func (c *下西里西亚Lower_silesiaCounty) BKrosno克罗斯诺() feud.Barony {
	return c.克罗斯诺Krosno
}
    
func (c *下西里西亚Lower_silesiaCounty) BNowogrod诺沃格鲁德() feud.Barony {
	return c.诺沃格鲁德Nowogrod
}
    
func (c *下西里西亚Lower_silesiaCounty) BSwiebozin希维博津() feud.Barony {
	return c.希维博津Swiebozin
}
    
func (c *下西里西亚Lower_silesiaCounty) BZagan扎甘() feud.Barony {
	return c.扎甘Zagan
}
    
func (c *下西里西亚Lower_silesiaCounty) BZielonagora绿山城() feud.Barony {
	return c.绿山城Zielonagora
}
    
var CLower_silesia下西里西亚 Lower_silesiaCounty = &下西里西亚Lower_silesiaCounty{}

func init() {
	f := CLower_silesia下西里西亚.(*下西里西亚Lower_silesiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "434",
		Title:     "lower_silesia",
		TitleName: "下西里西亚",
		TitleCode: "c_lower_silesia",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥得河畔比托姆Bytomodrz = BBytomodrz奥得河畔比托姆
	f.奥得河畔比托姆Bytomodrz.SetParent(f)
	
	f.格沃古夫Glogow = BGlogow格沃古夫
	f.格沃古夫Glogow.SetParent(f)
	
	f.克罗斯诺Krosno = BKrosno克罗斯诺
	f.克罗斯诺Krosno.SetParent(f)
	
	f.诺沃格鲁德Nowogrod = BNowogrod诺沃格鲁德
	f.诺沃格鲁德Nowogrod.SetParent(f)
	
	f.希维博津Swiebozin = BSwiebozin希维博津
	f.希维博津Swiebozin.SetParent(f)
	
	f.扎甘Zagan = BZagan扎甘
	f.扎甘Zagan.SetParent(f)
	
	f.绿山城Zielonagora = BZielonagora绿山城
	f.绿山城Zielonagora.SetParent(f)
	
}
