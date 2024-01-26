package usecase

import (
	"cloud_payments/internal/user"
	"encoding/json"
	"errors"
	"github.com/SevereCloud/vksdk/v2/api"
	"io"
	"net/http"
	"strconv"
)

func VkData(token, email string) (*user.SocialData, error) {

	var userData user.SocialData

	vk := api.NewVK(token)

	d, err := vk.UsersGet(map[string]interface{}{"fields": "photo_400_orig,photo_100,email"})

	if err != nil {
		return nil, err
	}

	if len(d) > 0 {
		u := d[0]

		socialID := strconv.Itoa(u.ID)

		userData = user.SocialData{
			ID:        socialID,
			Email:     email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Picture:   u.Photo100,
		}
	}

	if userData.ID == "" {
		return nil, errors.New("no id")
	}

	return &userData, nil

}

func GoogleData(token string) (*user.SocialData, error) {

	var userData user.SocialData

	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(bytes, &userData)

	if userData.ID == "" {
		return nil, errors.New("no id")
	}

	return &userData, nil

}
