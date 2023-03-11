package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LutskCounty interface {
    feud.County
    BDubno杜布诺() 	feud.Barony
    BJytomyr日托米尔() 	feud.Barony
    BLutsk卢茨克() 	feud.Barony
    BPeresopnytsia佩列索普尼察亚() 	feud.Barony
    BPochaiv波恰耶夫() 	feud.Barony
    BRivne罗夫诺() 	feud.Barony
    BShumsk舒姆西克() 	feud.Barony
}

type 卢茨克LutskCounty struct {
	feud.BaseCounty
	杜布诺Dubno 	feud.Barony
	日托米尔Jytomyr 	feud.Barony
	卢茨克Lutsk 	feud.Barony
	佩列索普尼察亚Peresopnytsia 	feud.Barony
	波恰耶夫Pochaiv 	feud.Barony
	罗夫诺Rivne 	feud.Barony
	舒姆西克Shumsk 	feud.Barony
}

func (c *卢茨克LutskCounty) BDubno杜布诺() feud.Barony {
	return c.杜布诺Dubno
}
    
func (c *卢茨克LutskCounty) BJytomyr日托米尔() feud.Barony {
	return c.日托米尔Jytomyr
}
    
func (c *卢茨克LutskCounty) BLutsk卢茨克() feud.Barony {
	return c.卢茨克Lutsk
}
    
func (c *卢茨克LutskCounty) BPeresopnytsia佩列索普尼察亚() feud.Barony {
	return c.佩列索普尼察亚Peresopnytsia
}
    
func (c *卢茨克LutskCounty) BPochaiv波恰耶夫() feud.Barony {
	return c.波恰耶夫Pochaiv
}
    
func (c *卢茨克LutskCounty) BRivne罗夫诺() feud.Barony {
	return c.罗夫诺Rivne
}
    
func (c *卢茨克LutskCounty) BShumsk舒姆西克() feud.Barony {
	return c.舒姆西克Shumsk
}
    
var CLutsk卢茨克 LutskCounty = &卢茨克LutskCounty{}

func init() {
	f := CLutsk卢茨克.(*卢茨克LutskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1633",
		Title:     "lutsk",
		TitleName: "卢茨克",
		TitleCode: "c_lutsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜布诺Dubno = BDubno杜布诺
	f.杜布诺Dubno.SetParent(f)
	
	f.日托米尔Jytomyr = BJytomyr日托米尔
	f.日托米尔Jytomyr.SetParent(f)
	
	f.卢茨克Lutsk = BLutsk卢茨克
	f.卢茨克Lutsk.SetParent(f)
	
	f.佩列索普尼察亚Peresopnytsia = BPeresopnytsia佩列索普尼察亚
	f.佩列索普尼察亚Peresopnytsia.SetParent(f)
	
	f.波恰耶夫Pochaiv = BPochaiv波恰耶夫
	f.波恰耶夫Pochaiv.SetParent(f)
	
	f.罗夫诺Rivne = BRivne罗夫诺
	f.罗夫诺Rivne.SetParent(f)
	
	f.舒姆西克Shumsk = BShumsk舒姆西克
	f.舒姆西克Shumsk.SetParent(f)
	
}
