package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Pereyaslavl_zalesskyCounty interface {
    feud.County
    BAleksandrov亚历山德罗夫() 	feud.Barony
    BKarabanovo卡拉巴诺沃() 	feud.Barony
    BKhmelniki赫梅利尼基() 	feud.Barony
    BKupanskoye库潘斯科耶() 	feud.Barony
    BPereyaslavlzalessky佩列亚斯拉夫尔() 	feud.Barony
    BSergiyevposad谢尔吉耶夫波萨德() 	feud.Barony
    BStrunino斯特鲁尼诺() 	feud.Barony
}

type 佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty struct {
	feud.BaseCounty
	亚历山德罗夫Aleksandrov 	feud.Barony
	卡拉巴诺沃Karabanovo 	feud.Barony
	赫梅利尼基Khmelniki 	feud.Barony
	库潘斯科耶Kupanskoye 	feud.Barony
	佩列亚斯拉夫尔Pereyaslavlzalessky 	feud.Barony
	谢尔吉耶夫波萨德Sergiyevposad 	feud.Barony
	斯特鲁尼诺Strunino 	feud.Barony
}

func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BAleksandrov亚历山德罗夫() feud.Barony {
	return c.亚历山德罗夫Aleksandrov
}
    
func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BKarabanovo卡拉巴诺沃() feud.Barony {
	return c.卡拉巴诺沃Karabanovo
}
    
func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BKhmelniki赫梅利尼基() feud.Barony {
	return c.赫梅利尼基Khmelniki
}
    
func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BKupanskoye库潘斯科耶() feud.Barony {
	return c.库潘斯科耶Kupanskoye
}
    
func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BPereyaslavlzalessky佩列亚斯拉夫尔() feud.Barony {
	return c.佩列亚斯拉夫尔Pereyaslavlzalessky
}
    
func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BSergiyevposad谢尔吉耶夫波萨德() feud.Barony {
	return c.谢尔吉耶夫波萨德Sergiyevposad
}
    
func (c *佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty) BStrunino斯特鲁尼诺() feud.Barony {
	return c.斯特鲁尼诺Strunino
}
    
var CPereyaslavl_zalessky佩列亚斯拉夫尔扎列斯基 Pereyaslavl_zalesskyCounty = &佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty{}

func init() {
	f := CPereyaslavl_zalessky佩列亚斯拉夫尔扎列斯基.(*佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalesskyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "573",
		Title:     "pereyaslavl_zalessky",
		TitleName: "佩列亚斯拉夫尔扎列斯基",
		TitleCode: "c_pereyaslavl_zalessky",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚历山德罗夫Aleksandrov = BAleksandrov亚历山德罗夫
	f.亚历山德罗夫Aleksandrov.SetParent(f)
	
	f.卡拉巴诺沃Karabanovo = BKarabanovo卡拉巴诺沃
	f.卡拉巴诺沃Karabanovo.SetParent(f)
	
	f.赫梅利尼基Khmelniki = BKhmelniki赫梅利尼基
	f.赫梅利尼基Khmelniki.SetParent(f)
	
	f.库潘斯科耶Kupanskoye = BKupanskoye库潘斯科耶
	f.库潘斯科耶Kupanskoye.SetParent(f)
	
	f.佩列亚斯拉夫尔Pereyaslavlzalessky = BPereyaslavlzalessky佩列亚斯拉夫尔
	f.佩列亚斯拉夫尔Pereyaslavlzalessky.SetParent(f)
	
	f.谢尔吉耶夫波萨德Sergiyevposad = BSergiyevposad谢尔吉耶夫波萨德
	f.谢尔吉耶夫波萨德Sergiyevposad.SetParent(f)
	
	f.斯特鲁尼诺Strunino = BStrunino斯特鲁尼诺
	f.斯特鲁尼诺Strunino.SetParent(f)
	
}
