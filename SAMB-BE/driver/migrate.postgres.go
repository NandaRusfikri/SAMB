package driver

import (
	"SAMB-BE/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.MasterProduct{},
		&entities.MasterWareHouse{},
		&entities.MasterCustomer{},
		&entities.MasterSupplier{},
		&entities.TrxPenerimaanBarangHeader{},
		&entities.TrxPenerimaanBarangDetail{},
		&entities.TrxPengeluaranBarangHeader{},
		&entities.TrxPengeluaranBarangDetail{},
	)

	if err != nil {
		log.Errorln("❌ Error Migrate ", err.Error())
		return err
	}

	WareHouse := []string{"Gudang A", "Gudang B", "Gudang C", "Gudang D"}
	if data := db.Find(&entities.MasterWareHouse{}); data.RowsAffected < 1 {
		for _, name := range WareHouse {
			db.Create(&entities.MasterWareHouse{
				WhsName: name,
			})
		}
		log.Println("✅ Seed MasterWareHouse inserted")
	}

	Products := []string{"SAMB-BEal ABC", "Indomie Goreng", "Qtela", "Promina"}
	if data := db.Find(&entities.MasterProduct{}); data.RowsAffected < 1 {
		for _, name := range Products {
			db.Create(&entities.MasterProduct{
				ProductName: name,
			})
		}
		log.Println("✅ Seed MasterProduct inserted")
	}
	Suppliers := []string{"Supplier A", "Supplier B", "Supplier C", "Supplier D"}
	if data := db.Find(&entities.MasterSupplier{}); data.RowsAffected < 1 {
		for _, name := range Suppliers {
			db.Create(&entities.MasterSupplier{
				SupplierName: name,
			})
		}
		log.Println("✅ Seed MasterSupplier inserted")
	}

	return err
}
