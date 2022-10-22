package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("projec-sa.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&User{},
		&Gender{},
		&Blood_type{},
		&Drug_Allergy{},
		&RIGHTS{},
		&Patient{},
	)

	//    add data 01 (add_patient)
	// //เปิดใช้งานเฉพาระตอนที่ในฐานข้อมูลว่างเปล่า เพราะจะทำการบันทึกซ้ำ
	// //ใส่ข้อมูลที่เป็นข้อมูลที่ไม่เปลี่ยนแปลง
	// //ตาราง หมู่เลือดs
	// database.Create(&Blood_type{Blood_Name: "A"})
	// database.Create(&Blood_type{Blood_Name: "B"})
	// database.Create(&Blood_type{Blood_Name: "O"})
	// database.Create(&Blood_type{Blood_Name: "AB"})
	// database.Create(&Blood_type{Blood_Name: "อื่น ๆ โปรดกรอกในช่องข้อมูลเพิ่มเติม"})
	// //ตาราง เพศ
	// database.Create(&Gender{Gender_Name: "ชาย"})
	// database.Create(&Gender{Gender_Name: "หญิง"})
	// //ตาราง ข้อมูลการแพ้ยา
	// database.Create(&Drug_Allergy{Drug_Name: "ไม่เคยแพ้ยา", Drug_Group: "-"})
	// database.Create(&Drug_Allergy{Drug_Name: "Paracetamol ", Drug_Group: "ยาระงับปวดและลดไข้"})
	// database.Create(&Drug_Allergy{Drug_Name: "Penicillin  ", Drug_Group: "ยาปฏิชีวนะ"})
	// database.Create(&Drug_Allergy{Drug_Name: "NSAIDs ", Drug_Group: "ยาแก้อักเสบ"})
	// database.Create(&Drug_Allergy{Drug_Name: "Carbamazepine ", Drug_Group: "ยากันชัก"})
	// database.Create(&Drug_Allergy{Drug_Name: "อื่น ๆ โปรดกรอกในช่องข้อมูลเพิ่มเติม", Drug_Group: "-"})
	// //ตาราง สิทธิ์การรักษาพยาบาล
	// database.Create(&RIGHTS{RIGHTS_Name: "ผู้ป่วยทั่วไปที่ไม่มีสิทธิ์การรักษาพยาบาล"})
	// database.Create(&RIGHTS{RIGHTS_Name: "สิทธิ์ข้าราชการ"})
	// database.Create(&RIGHTS{RIGHTS_Name: "สิทธิ์บัตรทอง"})
	// database.Create(&RIGHTS{RIGHTS_Name: "สิทธิ์ประกันสังคม"})
	// database.Create(&RIGHTS{RIGHTS_Name: "สิทธิ์องค์กรคู่สัญญา"})
	// database.Create(&RIGHTS{RIGHTS_Name: "ผู้ป่วยที่ถือบัตรประกันสุขภาพ"})
	// database.Create(&RIGHTS{RIGHTS_Name: "อื่น ๆ โปรดกรอกในช่องข้อมูลเพิ่มเติม"})

	db = database
}
