package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HendjanCounty interface {
    feud.County
    BArgan阿尔甘() 	feud.Barony
    BBandaremashoor班达尔马索尔() 	feud.Barony
    BBandarshahpur班达尔沙普尔() 	feud.Barony
    BBehbahan贝赫贝汉() 	feud.Barony
    BJayzan扎耶赞() 	feud.Barony
    BOmidiyeh奥米迪耶() 	feud.Barony
    BRamhormoz拉姆霍尔木兹() 	feud.Barony
    BSusa苏萨() 	feud.Barony
}

type 阿拉詹HendjanCounty struct {
	feud.BaseCounty
	阿尔甘Argan 	feud.Barony
	班达尔马索尔Bandaremashoor 	feud.Barony
	班达尔沙普尔Bandarshahpur 	feud.Barony
	贝赫贝汉Behbahan 	feud.Barony
	扎耶赞Jayzan 	feud.Barony
	奥米迪耶Omidiyeh 	feud.Barony
	拉姆霍尔木兹Ramhormoz 	feud.Barony
	苏萨Susa 	feud.Barony
}

func (c *阿拉詹HendjanCounty) BArgan阿尔甘() feud.Barony {
	return c.阿尔甘Argan
}
    
func (c *阿拉詹HendjanCounty) BBandaremashoor班达尔马索尔() feud.Barony {
	return c.班达尔马索尔Bandaremashoor
}
    
func (c *阿拉詹HendjanCounty) BBandarshahpur班达尔沙普尔() feud.Barony {
	return c.班达尔沙普尔Bandarshahpur
}
    
func (c *阿拉詹HendjanCounty) BBehbahan贝赫贝汉() feud.Barony {
	return c.贝赫贝汉Behbahan
}
    
func (c *阿拉詹HendjanCounty) BJayzan扎耶赞() feud.Barony {
	return c.扎耶赞Jayzan
}
    
func (c *阿拉詹HendjanCounty) BOmidiyeh奥米迪耶() feud.Barony {
	return c.奥米迪耶Omidiyeh
}
    
func (c *阿拉詹HendjanCounty) BRamhormoz拉姆霍尔木兹() feud.Barony {
	return c.拉姆霍尔木兹Ramhormoz
}
    
func (c *阿拉詹HendjanCounty) BSusa苏萨() feud.Barony {
	return c.苏萨Susa
}
    
var CHendjan阿拉詹 HendjanCounty = &阿拉詹HendjanCounty{}

func init() {
	f := CHendjan阿拉詹.(*阿拉詹HendjanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "645",
		Title:     "hendjan",
		TitleName: "阿拉詹",
		TitleCode: "c_hendjan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔甘Argan = BArgan阿尔甘
	f.阿尔甘Argan.SetParent(f)
	
	f.班达尔马索尔Bandaremashoor = BBandaremashoor班达尔马索尔
	f.班达尔马索尔Bandaremashoor.SetParent(f)
	
	f.班达尔沙普尔Bandarshahpur = BBandarshahpur班达尔沙普尔
	f.班达尔沙普尔Bandarshahpur.SetParent(f)
	
	f.贝赫贝汉Behbahan = BBehbahan贝赫贝汉
	f.贝赫贝汉Behbahan.SetParent(f)
	
	f.扎耶赞Jayzan = BJayzan扎耶赞
	f.扎耶赞Jayzan.SetParent(f)
	
	f.奥米迪耶Omidiyeh = BOmidiyeh奥米迪耶
	f.奥米迪耶Omidiyeh.SetParent(f)
	
	f.拉姆霍尔木兹Ramhormoz = BRamhormoz拉姆霍尔木兹
	f.拉姆霍尔木兹Ramhormoz.SetParent(f)
	
	f.苏萨Susa = BSusa苏萨
	f.苏萨Susa.SetParent(f)
	
}
