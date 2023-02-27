package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MinskCounty interface {
    feud.County
    BBarysaw鲍里索夫() 	feud.Barony
    BKoidanova科伊达诺瓦() 	feud.Barony
    BLahoysk洛戈伊斯克() 	feud.Barony
    BMaryina_horka马里纳戈尔卡() 	feud.Barony
    BMinsk明斯克() 	feud.Barony
    BNyasvizh涅斯维日() 	feud.Barony
    BZaslawye扎斯拉夫尔() 	feud.Barony
}

type 明斯克MinskCounty struct {
	feud.BaseCounty
	鲍里索夫Barysaw 	feud.Barony
	科伊达诺瓦Koidanova 	feud.Barony
	洛戈伊斯克Lahoysk 	feud.Barony
	马里纳戈尔卡Maryina_horka 	feud.Barony
	明斯克Minsk 	feud.Barony
	涅斯维日Nyasvizh 	feud.Barony
	扎斯拉夫尔Zaslawye 	feud.Barony
}

func (c *明斯克MinskCounty) BBarysaw鲍里索夫() feud.Barony {
	return c.鲍里索夫Barysaw
}
    
func (c *明斯克MinskCounty) BKoidanova科伊达诺瓦() feud.Barony {
	return c.科伊达诺瓦Koidanova
}
    
func (c *明斯克MinskCounty) BLahoysk洛戈伊斯克() feud.Barony {
	return c.洛戈伊斯克Lahoysk
}
    
func (c *明斯克MinskCounty) BMaryina_horka马里纳戈尔卡() feud.Barony {
	return c.马里纳戈尔卡Maryina_horka
}
    
func (c *明斯克MinskCounty) BMinsk明斯克() feud.Barony {
	return c.明斯克Minsk
}
    
func (c *明斯克MinskCounty) BNyasvizh涅斯维日() feud.Barony {
	return c.涅斯维日Nyasvizh
}
    
func (c *明斯克MinskCounty) BZaslawye扎斯拉夫尔() feud.Barony {
	return c.扎斯拉夫尔Zaslawye
}
    
var CMinsk明斯克 MinskCounty = &明斯克MinskCounty{}

func init() {
	f := CMinsk明斯克.(*明斯克MinskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "550",
		Title:     "minsk",
		TitleName: "明斯克",
		TitleCode: "c_minsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.鲍里索夫Barysaw = BBarysaw鲍里索夫
	f.鲍里索夫Barysaw.SetParent(f)
	
	f.科伊达诺瓦Koidanova = BKoidanova科伊达诺瓦
	f.科伊达诺瓦Koidanova.SetParent(f)
	
	f.洛戈伊斯克Lahoysk = BLahoysk洛戈伊斯克
	f.洛戈伊斯克Lahoysk.SetParent(f)
	
	f.马里纳戈尔卡Maryina_horka = BMaryina_horka马里纳戈尔卡
	f.马里纳戈尔卡Maryina_horka.SetParent(f)
	
	f.明斯克Minsk = BMinsk明斯克
	f.明斯克Minsk.SetParent(f)
	
	f.涅斯维日Nyasvizh = BNyasvizh涅斯维日
	f.涅斯维日Nyasvizh.SetParent(f)
	
	f.扎斯拉夫尔Zaslawye = BZaslawye扎斯拉夫尔
	f.扎斯拉夫尔Zaslawye.SetParent(f)
	
}
