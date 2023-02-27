package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QazanCounty interface {
    feud.County
    BBelyaevka别利亚耶夫卡() 	feud.Barony
    BChastye恰斯特耶() 	feud.Barony
    BKrmayak克尔马亚科() 	feud.Barony
    BOsa奥萨() 	feud.Barony
    BPakli帕克利() 	feud.Barony
    BPal帕利() 	feud.Barony
    BQazan喀山() 	feud.Barony
    BUsttuntor乌斯季_通托尔() 	feud.Barony
    BUymuzh维穆日() 	feud.Barony
}

type 喀山QazanCounty struct {
	feud.BaseCounty
	别利亚耶夫卡Belyaevka 	feud.Barony
	恰斯特耶Chastye 	feud.Barony
	克尔马亚科Krmayak 	feud.Barony
	奥萨Osa 	feud.Barony
	帕克利Pakli 	feud.Barony
	帕利Pal 	feud.Barony
	喀山Qazan 	feud.Barony
	乌斯季_通托尔Usttuntor 	feud.Barony
	维穆日Uymuzh 	feud.Barony
}

func (c *喀山QazanCounty) BBelyaevka别利亚耶夫卡() feud.Barony {
	return c.别利亚耶夫卡Belyaevka
}
    
func (c *喀山QazanCounty) BChastye恰斯特耶() feud.Barony {
	return c.恰斯特耶Chastye
}
    
func (c *喀山QazanCounty) BKrmayak克尔马亚科() feud.Barony {
	return c.克尔马亚科Krmayak
}
    
func (c *喀山QazanCounty) BOsa奥萨() feud.Barony {
	return c.奥萨Osa
}
    
func (c *喀山QazanCounty) BPakli帕克利() feud.Barony {
	return c.帕克利Pakli
}
    
func (c *喀山QazanCounty) BPal帕利() feud.Barony {
	return c.帕利Pal
}
    
func (c *喀山QazanCounty) BQazan喀山() feud.Barony {
	return c.喀山Qazan
}
    
func (c *喀山QazanCounty) BUsttuntor乌斯季_通托尔() feud.Barony {
	return c.乌斯季_通托尔Usttuntor
}
    
func (c *喀山QazanCounty) BUymuzh维穆日() feud.Barony {
	return c.维穆日Uymuzh
}
    
var CQazan喀山 QazanCounty = &喀山QazanCounty{}

func init() {
	f := CQazan喀山.(*喀山QazanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "611",
		Title:     "qazan",
		TitleName: "喀山",
		TitleCode: "c_qazan",
		Baronies:  map[string]feud.Barony{},
	}

	f.别利亚耶夫卡Belyaevka = BBelyaevka别利亚耶夫卡
	f.别利亚耶夫卡Belyaevka.SetParent(f)
	
	f.恰斯特耶Chastye = BChastye恰斯特耶
	f.恰斯特耶Chastye.SetParent(f)
	
	f.克尔马亚科Krmayak = BKrmayak克尔马亚科
	f.克尔马亚科Krmayak.SetParent(f)
	
	f.奥萨Osa = BOsa奥萨
	f.奥萨Osa.SetParent(f)
	
	f.帕克利Pakli = BPakli帕克利
	f.帕克利Pakli.SetParent(f)
	
	f.帕利Pal = BPal帕利
	f.帕利Pal.SetParent(f)
	
	f.喀山Qazan = BQazan喀山
	f.喀山Qazan.SetParent(f)
	
	f.乌斯季_通托尔Usttuntor = BUsttuntor乌斯季_通托尔
	f.乌斯季_通托尔Usttuntor.SetParent(f)
	
	f.维穆日Uymuzh = BUymuzh维穆日
	f.维穆日Uymuzh.SetParent(f)
	
}
