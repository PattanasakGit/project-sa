## Update
1. ได้ทำการรวม controller ทั้งหมดไว้เป็นไฟล์เดียว --> backend/controller/PatientController.go เอาไว้กรณีรวมงาน
2. ได้ทำการแก้ไขในส่วนของ schema ทำการเพิ่ม `gorm:"references:id"` (ที่จั๊ดลืมไว้) ตอนนี้เพิ่มเรียบร้อยแล้ว
## สถานะ ของโปรแกรม
1. รันได้
2. saveได้ แต่ยังไม่มีส่วนของ User
## สิ่งที่ต้องแก้ไข
1. ปรับ UI ลบปุ่ม ย้อนกลับออก
2. ทำการจัดเก็บส่วนของ User ลงตารางหลัก
