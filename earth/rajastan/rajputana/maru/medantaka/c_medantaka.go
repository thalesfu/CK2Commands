package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedantakaCounty interface {
	feud.County
	BMedantaka迷檀多迦() feud.Barony
	BOudh奥德() feud.Barony
	BPachaur波招罗() feud.Barony
	BPakal波伽() feud.Barony
	BPhalarddhi颇罗梨地() feud.Barony
	BPurnatallakapura富楼那呾罗迦补罗() feud.Barony
	BSanjoo桑耆() feud.Barony
}

type 迷檀多迦MedantakaCounty struct {
	feud.BaseCounty
	迷檀多迦Medantaka            feud.Barony
	奥德Oudh                   feud.Barony
	波招罗Pachaur               feud.Barony
	波伽Pakal                  feud.Barony
	颇罗梨地Phalarddhi           feud.Barony
	富楼那呾罗迦补罗Purnatallakapura feud.Barony
	桑耆Sanjoo                 feud.Barony
}

func (c *迷檀多迦MedantakaCounty) BMedantaka迷檀多迦() feud.Barony {
	return c.迷檀多迦Medantaka
}

func (c *迷檀多迦MedantakaCounty) BOudh奥德() feud.Barony {
	return c.奥德Oudh
}

func (c *迷檀多迦MedantakaCounty) BPachaur波招罗() feud.Barony {
	return c.波招罗Pachaur
}

func (c *迷檀多迦MedantakaCounty) BPakal波伽() feud.Barony {
	return c.波伽Pakal
}

func (c *迷檀多迦MedantakaCounty) BPhalarddhi颇罗梨地() feud.Barony {
	return c.颇罗梨地Phalarddhi
}

func (c *迷檀多迦MedantakaCounty) BPurnatallakapura富楼那呾罗迦补罗() feud.Barony {
	return c.富楼那呾罗迦补罗Purnatallakapura
}

func (c *迷檀多迦MedantakaCounty) BSanjoo桑耆() feud.Barony {
	return c.桑耆Sanjoo
}

var CMedantaka迷檀多迦 MedantakaCounty = &迷檀多迦MedantakaCounty{}

func init() {
	f := CMedantaka迷檀多迦.(*迷檀多迦MedantakaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1176",
		Title:     "medantaka",
		TitleName: "迷檀多迦",
		TitleCode: "c_medantaka",
		Baronies:  map[string]feud.Barony{},
	}

	f.迷檀多迦Medantaka = BMedantaka迷檀多迦
	f.迷檀多迦Medantaka.SetParent(f)

	f.奥德Oudh = BOudh奥德
	f.奥德Oudh.SetParent(f)

	f.波招罗Pachaur = BPachaur波招罗
	f.波招罗Pachaur.SetParent(f)

	f.波伽Pakal = BPakal波伽
	f.波伽Pakal.SetParent(f)

	f.颇罗梨地Phalarddhi = BPhalarddhi颇罗梨地
	f.颇罗梨地Phalarddhi.SetParent(f)

	f.富楼那呾罗迦补罗Purnatallakapura = BPurnatallakapura富楼那呾罗迦补罗
	f.富楼那呾罗迦补罗Purnatallakapura.SetParent(f)

	f.桑耆Sanjoo = BSanjoo桑耆
	f.桑耆Sanjoo.SetParent(f)

}
