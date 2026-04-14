package main

/*
import (
	"BE_Go/config"
	"BE_Go/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}
//
func main() {
	config.ConnectDB()
	db := config.DB

	password := hashPassword("tes123")
	fmt.Println("Starting seed...")

	// ================= ADMIN =================
	var adminCount int64
	db.Model(&models.Admin{}).Where("username = ?", "admin").Count(&adminCount)
	if adminCount == 0 {
		admin := models.Admin{
			Name:         "System Admin",
			Username:     "admin",
			PasswordHash: password,
		}
		db.Create(&admin)
		fmt.Println("✅ Admin created")
	} else {
		fmt.Println("⚠️  Admin already exists")
	}

	// ================= HOUSES =================
	housesData := []struct {
		name        string
		description string
		logoURL     string
	}{
		{
			"Al-Ghuraab",
			"Al-Ghuraab (الغراب) — Inspired by the crow mentioned in the Qur'an. Represents learning through observation, humility, and moral awareness.",
			"https://res.cloudinary.com/dntujhjkw/image/upload/v1770108364/Al-Ghuraab_szxjhl.png",
		},
		{
			"An-Nahl",
			"An-Nahl (النحل) — Inspired by the bee mentioned in the Qur'an. Symbolizes productivity, order, obedience, and service to others.",
			"https://res.cloudinary.com/dntujhjkw/image/upload/v1770108374/An-Nahl_pelgou.png",
		},
		{
			"An-Nun",
			"An-Nun (النون) — Inspired by the great fish associated with Prophet Yunus. Represents patience, repentance, resilience, and self-reflection.",
			"https://res.cloudinary.com/dntujhjkw/image/upload/v1770108369/An-Nun_erm5nb.png",
		},
		{
			"Al-Adiyat",
			"Al-Adiyat (العاديات) — Inspired by the charging horses mentioned in the Qur'an. Symbolizes discipline, loyalty, determination, and relentless effort.",
			"https://res.cloudinary.com/dntujhjkw/image/upload/v1770108368/Al-Adiyat_l5h9eh.png",
		},
		{
			"Al-Hudhud",
			"Al-Hudhud (الهدهد) — Inspired by the hoopoe bird mentioned in the Qur'an. Represents intelligence, communication, courage, and responsibility.",
			"https://res.cloudinary.com/dntujhjkw/image/upload/v1770108363/Al-HudHud_oblb0w.png",
		},
		{
			"An-Naml",
			"An-Naml (النمل) — Inspired by the ants mentioned in the Qur'an. Symbolizes teamwork, awareness, humility, and care for the community.",
			"https://res.cloudinary.com/dntujhjkw/image/upload/v1770108363/An-Naml_hqfrmx.png",
		},
	}

	houses := map[string]models.House{}
	for _, h := range housesData {
		var house models.House
		result := db.Where("name = ?", h.name).First(&house)
		if result.Error != nil {
			house = models.House{
				Name:        h.name,
				Description: h.description,
				HousePoints: 0,
				LogoURL:     h.logoURL,
			}
			db.Create(&house)
			fmt.Printf("✅ House created: %s\n", h.name)
		} else {
			fmt.Printf("⚠️  House already exists: %s\n", h.name)
		}
		houses[h.name] = house
	}

	// ================= CAPTAINS =================
	captainsData := []struct {
		name      string
		username  string
		houseName string
	}{
		{"Captain Ghuraab", "ghuraab", "Al-Ghuraab"},
		{"Captain Nahl", "nahl", "An-Nahl"},
		{"Captain Nun", "nun", "An-Nun"},
		{"Captain Adiyat", "adiyat", "Al-Adiyat"},
		{"Captain Hudhud", "hudhud", "Al-Hudhud"},
		{"Captain Naml", "naml", "An-Naml"},
	}

	for _, c := range captainsData {
		var count int64
		db.Model(&models.Captain{}).Where("username = ?", c.username).Count(&count)
		if count == 0 {
			house := houses[c.houseName]
			captain := models.Captain{
				Name:         c.name,
				Username:     c.username,
				PasswordHash: password,
				HouseID:      int(house.ID),
			}
			db.Create(&captain)
			fmt.Printf("✅ Captain created: %s\n", c.name)
		} else {
			fmt.Printf("⚠️  Captain already exists: %s\n", c.username)
		}
	}

	// ================= MEMBERS =================
	membersData := []struct {
		name      string
		role      string
		houseName string
	}{
		{"Ahmad", "Member", "Al-Ghuraab"},
		{"Fatimah", "Member", "Al-Ghuraab"},
		{"Ali", "Member", "An-Nahl"},
		{"Amina", "Member", "An-Nahl"},
		{"Umar", "Member", "An-Nun"},
		{"Khadijah", "Member", "An-Nun"},
		{"Hasan", "Member", "Al-Adiyat"},
		{"Husain", "Member", "Al-Adiyat"},
		{"Bilal", "Member", "Al-Hudhud"},
		{"Zainab", "Member", "Al-Hudhud"},
		{"Yasir", "Member", "An-Naml"},
		{"Maryam", "Member", "An-Naml"},
	}

	for _, m := range membersData {
		house := houses[m.houseName]
		var count int64
		db.Model(&models.Member{}).Where("name = ? AND house_id = ?", m.name, house.ID).Count(&count)
		if count == 0 {
			member := models.Member{
				Name:    m.name,
				Role:    m.role,
				HouseID: int(house.ID),
			}
			db.Create(&member)
			fmt.Printf("✅ Member created: %s (%s)\n", m.name, m.houseName)
		} else {
			fmt.Printf("⚠️  Member already exists: %s\n", m.name)
		}
	}

	// ================= ANNOUNCEMENTS =================
	announcementsData := []struct {
		title   string
		content string
	}{
		{"Welcome Announcement", "Welcome to the new house season!"},
		{"Training Session", "House training will begin this Friday."},
		{"Team Reminder", "Remember to wear house shirts every Monday."},
	}

	var captains []models.Captain
	db.Find(&captains)

	for _, captain := range captains {
		for _, a := range announcementsData {
			var count int64
			db.Model(&models.Announcement{}).
				Where("title = ? AND captain_id = ?", a.title, captain.ID).
				Count(&count)
			if count == 0 {
				captainID := captain.ID
				announcement := models.Announcement{
					Title:     a.title,
					Content:   a.content,
					HouseID:   captain.HouseID,
					CaptainID: &captainID,
					CreatedAt: time.Now(),
				}
				db.Create(&announcement)
				fmt.Printf("✅ Announcement created: '%s' for captain %s\n", a.title, captain.Name)
			}
		}
	}

	fmt.Println("\n🌱 Seeding completed successfully!")
}
*/
