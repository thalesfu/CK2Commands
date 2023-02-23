package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PratishthanaCounty interface {
	feud.County
	BAhmednagar艾哈迈德纳格尔() feud.Barony
	BBhingar平格尔() feud.Barony
	BChambargonda占跋罗郡荼() feud.Barony
	BHarishchandragad诃利施旃陀罗伽荼() feud.Barony
	BPatkapur帕提加普尔() feud.Barony
	BPochannapet菩旃那毗提() feud.Barony
	BPratishthana波罗提瑟吒那() feud.Barony
}

type 波罗提瑟吒那PratishthanaCounty struct {
	feud.BaseCounty
	艾哈迈德纳格尔Ahmednagar        feud.Barony
	平格尔Bhingar               feud.Barony
	占跋罗郡荼Chambargonda        feud.Barony
	诃利施旃陀罗伽荼Harishchandragad feud.Barony
	帕提加普尔Patkapur            feud.Barony
	菩旃那毗提Pochannapet         feud.Barony
	波罗提瑟吒那Pratishthana       feud.Barony
}

func (c *波罗提瑟吒那PratishthanaCounty) BAhmednagar艾哈迈德纳格尔() feud.Barony {
	return c.艾哈迈德纳格尔Ahmednagar
}

func (c *波罗提瑟吒那PratishthanaCounty) BBhingar平格尔() feud.Barony {
	return c.平格尔Bhingar
}

func (c *波罗提瑟吒那PratishthanaCounty) BChambargonda占跋罗郡荼() feud.Barony {
	return c.占跋罗郡荼Chambargonda
}

func (c *波罗提瑟吒那PratishthanaCounty) BHarishchandragad诃利施旃陀罗伽荼() feud.Barony {
	return c.诃利施旃陀罗伽荼Harishchandragad
}

func (c *波罗提瑟吒那PratishthanaCounty) BPatkapur帕提加普尔() feud.Barony {
	return c.帕提加普尔Patkapur
}

func (c *波罗提瑟吒那PratishthanaCounty) BPochannapet菩旃那毗提() feud.Barony {
	return c.菩旃那毗提Pochannapet
}

func (c *波罗提瑟吒那PratishthanaCounty) BPratishthana波罗提瑟吒那() feud.Barony {
	return c.波罗提瑟吒那Pratishthana
}

var CPratishthana波罗提瑟吒那 PratishthanaCounty = &波罗提瑟吒那PratishthanaCounty{}

func init() {
	f := CPratishthana波罗提瑟吒那.(*波罗提瑟吒那PratishthanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1142",
		Title:     "pratishthana",
		TitleName: "波罗提瑟吒那",
		TitleCode: "c_pratishthana",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾哈迈德纳格尔Ahmednagar = BAhmednagar艾哈迈德纳格尔
	f.艾哈迈德纳格尔Ahmednagar.SetParent(f)

	f.平格尔Bhingar = BBhingar平格尔
	f.平格尔Bhingar.SetParent(f)

	f.占跋罗郡荼Chambargonda = BChambargonda占跋罗郡荼
	f.占跋罗郡荼Chambargonda.SetParent(f)

	f.诃利施旃陀罗伽荼Harishchandragad = BHarishchandragad诃利施旃陀罗伽荼
	f.诃利施旃陀罗伽荼Harishchandragad.SetParent(f)

	f.帕提加普尔Patkapur = BPatkapur帕提加普尔
	f.帕提加普尔Patkapur.SetParent(f)

	f.菩旃那毗提Pochannapet = BPochannapet菩旃那毗提
	f.菩旃那毗提Pochannapet.SetParent(f)

	f.波罗提瑟吒那Pratishthana = BPratishthana波罗提瑟吒那
	f.波罗提瑟吒那Pratishthana.SetParent(f)

}
