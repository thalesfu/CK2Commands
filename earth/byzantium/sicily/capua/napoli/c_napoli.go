package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NapoliCounty interface {
    feud.County
    BAfragola阿夫拉戈拉() 	feud.Barony
    BAversa阿韦尔萨() 	feud.Barony
    BCumae库迈() 	feud.Barony
    BIschia伊斯基亚() 	feud.Barony
    BNapoli那波利() 	feud.Barony
    BPortici波尔蒂奇() 	feud.Barony
    BPozzuoli波佐利() 	feud.Barony
    BTurris_octava图里斯_奥克塔瓦() 	feud.Barony
}

type 那波利NapoliCounty struct {
	feud.BaseCounty
	阿夫拉戈拉Afragola 	feud.Barony
	阿韦尔萨Aversa 	feud.Barony
	库迈Cumae 	feud.Barony
	伊斯基亚Ischia 	feud.Barony
	那波利Napoli 	feud.Barony
	波尔蒂奇Portici 	feud.Barony
	波佐利Pozzuoli 	feud.Barony
	图里斯_奥克塔瓦Turris_octava 	feud.Barony
}

func (c *那波利NapoliCounty) BAfragola阿夫拉戈拉() feud.Barony {
	return c.阿夫拉戈拉Afragola
}
    
func (c *那波利NapoliCounty) BAversa阿韦尔萨() feud.Barony {
	return c.阿韦尔萨Aversa
}
    
func (c *那波利NapoliCounty) BCumae库迈() feud.Barony {
	return c.库迈Cumae
}
    
func (c *那波利NapoliCounty) BIschia伊斯基亚() feud.Barony {
	return c.伊斯基亚Ischia
}
    
func (c *那波利NapoliCounty) BNapoli那波利() feud.Barony {
	return c.那波利Napoli
}
    
func (c *那波利NapoliCounty) BPortici波尔蒂奇() feud.Barony {
	return c.波尔蒂奇Portici
}
    
func (c *那波利NapoliCounty) BPozzuoli波佐利() feud.Barony {
	return c.波佐利Pozzuoli
}
    
func (c *那波利NapoliCounty) BTurris_octava图里斯_奥克塔瓦() feud.Barony {
	return c.图里斯_奥克塔瓦Turris_octava
}
    
var CNapoli那波利 NapoliCounty = &那波利NapoliCounty{}

func init() {
	f := CNapoli那波利.(*那波利NapoliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "334",
		Title:     "napoli",
		TitleName: "那波利",
		TitleCode: "c_napoli",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫拉戈拉Afragola = BAfragola阿夫拉戈拉
	f.阿夫拉戈拉Afragola.SetParent(f)
	
	f.阿韦尔萨Aversa = BAversa阿韦尔萨
	f.阿韦尔萨Aversa.SetParent(f)
	
	f.库迈Cumae = BCumae库迈
	f.库迈Cumae.SetParent(f)
	
	f.伊斯基亚Ischia = BIschia伊斯基亚
	f.伊斯基亚Ischia.SetParent(f)
	
	f.那波利Napoli = BNapoli那波利
	f.那波利Napoli.SetParent(f)
	
	f.波尔蒂奇Portici = BPortici波尔蒂奇
	f.波尔蒂奇Portici.SetParent(f)
	
	f.波佐利Pozzuoli = BPozzuoli波佐利
	f.波佐利Pozzuoli.SetParent(f)
	
	f.图里斯_奥克塔瓦Turris_octava = BTurris_octava图里斯_奥克塔瓦
	f.图里斯_奥克塔瓦Turris_octava.SetParent(f)
	
}
