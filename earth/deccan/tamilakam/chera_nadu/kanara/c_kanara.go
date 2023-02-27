package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanaraCounty interface {
    feud.County
    BAruvankadu阿罗槃迦株() 	feud.Barony
    BDharmasthala达摩萨他罗() 	feud.Barony
    BKadarika伽陀利迦() 	feud.Barony
    BKasargod伽萨罗() 	feud.Barony
    BMangalur莽葛奴儿() 	feud.Barony
    BUdayavara优陀耶伐罗() 	feud.Barony
    BUdupi尤杜毗() 	feud.Barony
}

type 伽那罗KanaraCounty struct {
	feud.BaseCounty
	阿罗槃迦株Aruvankadu 	feud.Barony
	达摩萨他罗Dharmasthala 	feud.Barony
	伽陀利迦Kadarika 	feud.Barony
	伽萨罗Kasargod 	feud.Barony
	莽葛奴儿Mangalur 	feud.Barony
	优陀耶伐罗Udayavara 	feud.Barony
	尤杜毗Udupi 	feud.Barony
}

func (c *伽那罗KanaraCounty) BAruvankadu阿罗槃迦株() feud.Barony {
	return c.阿罗槃迦株Aruvankadu
}
    
func (c *伽那罗KanaraCounty) BDharmasthala达摩萨他罗() feud.Barony {
	return c.达摩萨他罗Dharmasthala
}
    
func (c *伽那罗KanaraCounty) BKadarika伽陀利迦() feud.Barony {
	return c.伽陀利迦Kadarika
}
    
func (c *伽那罗KanaraCounty) BKasargod伽萨罗() feud.Barony {
	return c.伽萨罗Kasargod
}
    
func (c *伽那罗KanaraCounty) BMangalur莽葛奴儿() feud.Barony {
	return c.莽葛奴儿Mangalur
}
    
func (c *伽那罗KanaraCounty) BUdayavara优陀耶伐罗() feud.Barony {
	return c.优陀耶伐罗Udayavara
}
    
func (c *伽那罗KanaraCounty) BUdupi尤杜毗() feud.Barony {
	return c.尤杜毗Udupi
}
    
var CKanara伽那罗 KanaraCounty = &伽那罗KanaraCounty{}

func init() {
	f := CKanara伽那罗.(*伽那罗KanaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1200",
		Title:     "kanara",
		TitleName: "伽那罗",
		TitleCode: "c_kanara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿罗槃迦株Aruvankadu = BAruvankadu阿罗槃迦株
	f.阿罗槃迦株Aruvankadu.SetParent(f)
	
	f.达摩萨他罗Dharmasthala = BDharmasthala达摩萨他罗
	f.达摩萨他罗Dharmasthala.SetParent(f)
	
	f.伽陀利迦Kadarika = BKadarika伽陀利迦
	f.伽陀利迦Kadarika.SetParent(f)
	
	f.伽萨罗Kasargod = BKasargod伽萨罗
	f.伽萨罗Kasargod.SetParent(f)
	
	f.莽葛奴儿Mangalur = BMangalur莽葛奴儿
	f.莽葛奴儿Mangalur.SetParent(f)
	
	f.优陀耶伐罗Udayavara = BUdayavara优陀耶伐罗
	f.优陀耶伐罗Udayavara.SetParent(f)
	
	f.尤杜毗Udupi = BUdupi尤杜毗
	f.尤杜毗Udupi.SetParent(f)
	
}
