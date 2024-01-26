package repo

import (
	"cloud_payments/internal/models"
	"cloud_payments/internal/user"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type UserRepo struct {
	Db *gorm.DB
}

func (r *UserRepo) CheckUserBan(u *models.User, req *http.Request) ([]models.UserBan, error) {

	//t := time.Now()

	ip := req.Header.Get("X-Forwarded-For")
	ipb := strings.Split(ip, ".")
	if len(ipb) == 4 {
		ipb[3] = "*"
	}
	ipMask := strings.Join(ipb, ".")

	var bans []models.UserBan
	r.Db.Model(&models.UserBan{}).Where("(user_id = ? OR ip LIKE ? OR ip LIKE ?) "+
		"AND banned_until > CURRENT_TIMESTAMP", u.ID, ip, ipMask).Find(&bans)

	if len(bans) > 0 {
		return bans, nil
	}

	return nil, errors.New("not found")
}

func (r *UserRepo) NewUserVisitLog(u *models.User, req *http.Request) error {

	tx := r.Db.Create(&models.UserAuthorization{
		IP:     req.Header.Get("X-Forwarded-For"),
		UA:     req.UserAgent(),
		UserID: u.ID,
	})

	if tx.Error != nil {
		return tx.Error
	}

	return nil

}

func (r *UserRepo) GetUsers(limit, offset int) []models.User {

	u := []models.User{}

	r.Db.Model(&models.User{}).
		Preload("UserBans").
		Preload("UserAuthorizations").
		Order("id desc").
		Limit(limit).
		Offset(offset).
		Find(&u)

	return u
}

func (r *UserRepo) GetUserByID(id int) (*models.User, error) {

	var u models.User
	tx := r.Db.Model(&models.User{}).Where("id = ?", id).First(&u)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &u, nil
}

func (r *UserRepo) RegisterSocial(data *user.SocialData, typ string) (*models.User, error) {

	u := models.User{}

	r.Db.Model(&models.User{}).Where("type = ? and social_id = ?", typ, data.ID).First(&u)

	if u.ID > 0 {
		return &u, nil
	}

	// Иначе создадим

	u = models.User{
		Type:      typ,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Image:     data.Picture,
		SocialID:  data.ID,
		Email:     data.Email,
	}

	tx := r.Db.Create(&u)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &u, nil
}

func (r *UserRepo) RegisterAnonimous(name, typ string) (*models.User, error) {

	u := models.User{
		Type:      typ,
		FirstName: name,
	}

	tx := r.Db.Create(&u)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &u, nil

}
