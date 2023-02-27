package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BorovichiCounty interface {
    feud.County
    BBorovichi博罗维奇() 	feud.Barony
    BKhoromy霍罗梅() 	feud.Barony
    BKirichi基里奇() 	feud.Barony
    BLyubytino柳贝季诺() 	feud.Barony
    BOkulovka奥库洛夫卡() 	feud.Barony
    BPchevzha普切夫扎() 	feud.Barony
    BVodogon沃多贡() 	feud.Barony
}

type 博罗维奇BorovichiCounty struct {
	feud.BaseCounty
	博罗维奇Borovichi 	feud.Barony
	霍罗梅Khoromy 	feud.Barony
	基里奇Kirichi 	feud.Barony
	柳贝季诺Lyubytino 	feud.Barony
	奥库洛夫卡Okulovka 	feud.Barony
	普切夫扎Pchevzha 	feud.Barony
	沃多贡Vodogon 	feud.Barony
}

func (c *博罗维奇BorovichiCounty) BBorovichi博罗维奇() feud.Barony {
	return c.博罗维奇Borovichi
}
    
func (c *博罗维奇BorovichiCounty) BKhoromy霍罗梅() feud.Barony {
	return c.霍罗梅Khoromy
}
    
func (c *博罗维奇BorovichiCounty) BKirichi基里奇() feud.Barony {
	return c.基里奇Kirichi
}
    
func (c *博罗维奇BorovichiCounty) BLyubytino柳贝季诺() feud.Barony {
	return c.柳贝季诺Lyubytino
}
    
func (c *博罗维奇BorovichiCounty) BOkulovka奥库洛夫卡() feud.Barony {
	return c.奥库洛夫卡Okulovka
}
    
func (c *博罗维奇BorovichiCounty) BPchevzha普切夫扎() feud.Barony {
	return c.普切夫扎Pchevzha
}
    
func (c *博罗维奇BorovichiCounty) BVodogon沃多贡() feud.Barony {
	return c.沃多贡Vodogon
}
    
var CBorovichi博罗维奇 BorovichiCounty = &博罗维奇BorovichiCounty{}

func init() {
	f := CBorovichi博罗维奇.(*博罗维奇BorovichiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1652",
		Title:     "borovichi",
		TitleName: "博罗维奇",
		TitleCode: "c_borovichi",
		Baronies:  map[string]feud.Barony{},
	}

	f.博罗维奇Borovichi = BBorovichi博罗维奇
	f.博罗维奇Borovichi.SetParent(f)
	
	f.霍罗梅Khoromy = BKhoromy霍罗梅
	f.霍罗梅Khoromy.SetParent(f)
	
	f.基里奇Kirichi = BKirichi基里奇
	f.基里奇Kirichi.SetParent(f)
	
	f.柳贝季诺Lyubytino = BLyubytino柳贝季诺
	f.柳贝季诺Lyubytino.SetParent(f)
	
	f.奥库洛夫卡Okulovka = BOkulovka奥库洛夫卡
	f.奥库洛夫卡Okulovka.SetParent(f)
	
	f.普切夫扎Pchevzha = BPchevzha普切夫扎
	f.普切夫扎Pchevzha.SetParent(f)
	
	f.沃多贡Vodogon = BVodogon沃多贡
	f.沃多贡Vodogon.SetParent(f)
	
}
