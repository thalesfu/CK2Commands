package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalpiCounty interface {
    feud.County
    BGalleborgao伽勒菩乔() 	feud.Barony
    BJalaun耶罗郁那() 	feud.Barony
    BKalpi伽罗毗() 	feud.Barony
    BKonch军遮() 	feud.Barony
    BKurara俱罗罗() 	feud.Barony
    BLahar拉诃罗() 	feud.Barony
    BOirai乌伊罗伊() 	feud.Barony
}

type 伽罗毗KalpiCounty struct {
	feud.BaseCounty
	伽勒菩乔Galleborgao 	feud.Barony
	耶罗郁那Jalaun 	feud.Barony
	伽罗毗Kalpi 	feud.Barony
	军遮Konch 	feud.Barony
	俱罗罗Kurara 	feud.Barony
	拉诃罗Lahar 	feud.Barony
	乌伊罗伊Oirai 	feud.Barony
}

func (c *伽罗毗KalpiCounty) BGalleborgao伽勒菩乔() feud.Barony {
	return c.伽勒菩乔Galleborgao
}
    
func (c *伽罗毗KalpiCounty) BJalaun耶罗郁那() feud.Barony {
	return c.耶罗郁那Jalaun
}
    
func (c *伽罗毗KalpiCounty) BKalpi伽罗毗() feud.Barony {
	return c.伽罗毗Kalpi
}
    
func (c *伽罗毗KalpiCounty) BKonch军遮() feud.Barony {
	return c.军遮Konch
}
    
func (c *伽罗毗KalpiCounty) BKurara俱罗罗() feud.Barony {
	return c.俱罗罗Kurara
}
    
func (c *伽罗毗KalpiCounty) BLahar拉诃罗() feud.Barony {
	return c.拉诃罗Lahar
}
    
func (c *伽罗毗KalpiCounty) BOirai乌伊罗伊() feud.Barony {
	return c.乌伊罗伊Oirai
}
    
var CKalpi伽罗毗 KalpiCounty = &伽罗毗KalpiCounty{}

func init() {
	f := CKalpi伽罗毗.(*伽罗毗KalpiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1169",
		Title:     "kalpi",
		TitleName: "伽罗毗",
		TitleCode: "c_kalpi",
		Baronies:  map[string]feud.Barony{},
	}

	f.伽勒菩乔Galleborgao = BGalleborgao伽勒菩乔
	f.伽勒菩乔Galleborgao.SetParent(f)
	
	f.耶罗郁那Jalaun = BJalaun耶罗郁那
	f.耶罗郁那Jalaun.SetParent(f)
	
	f.伽罗毗Kalpi = BKalpi伽罗毗
	f.伽罗毗Kalpi.SetParent(f)
	
	f.军遮Konch = BKonch军遮
	f.军遮Konch.SetParent(f)
	
	f.俱罗罗Kurara = BKurara俱罗罗
	f.俱罗罗Kurara.SetParent(f)
	
	f.拉诃罗Lahar = BLahar拉诃罗
	f.拉诃罗Lahar.SetParent(f)
	
	f.乌伊罗伊Oirai = BOirai乌伊罗伊
	f.乌伊罗伊Oirai.SetParent(f)
	
}
