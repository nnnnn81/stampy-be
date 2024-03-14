package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nnnnn81/stampy-be/db"
	"github.com/nnnnn81/stampy-be/model"
	"gorm.io/gorm"
)

func CardShow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)
	var cards []model.Stampcard
	if err := db.DB.Where("created_by = ?", userid).Find(&cards).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"cards": []model.Stampcard{},
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		var responseData []echo.Map
		for _, card := range cards {
			// createdUser取得
			var createduser model.User
			if err := db.DB.Where("id = ?", card.CreatedBy).First(&createduser).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// return 404
					return c.JSON(http.StatusNotFound, echo.Map{
						"message": "User Not Found",
					})

				} else {
					// return 500
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"message": "Database Error: " + err.Error(),
					})
				}
			}

			// joinedUser取得
			var joineduser model.User
			if err := db.DB.Where("email = ?", card.JoinedUser).First(&joineduser).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// return 404
					return c.JSON(http.StatusNotFound, echo.Map{
						"message": "User Not Found",
					})

				} else {
					// return 500
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"message": "Database Error: " + err.Error(),
					})
				}
			}
			var omitedcreateduser model.OmitUser
			omitedcreateduser.Id = createduser.Id
			omitedcreateduser.Email = createduser.Email
			omitedcreateduser.Username = createduser.Username
			omitedcreateduser.AvaterUrl = createduser.AvaterUrl
			var omitedjoineduser model.OmitUser
			omitedjoineduser.Id = joineduser.Id
			omitedjoineduser.Email = joineduser.Email
			omitedjoineduser.Username = joineduser.Username
			omitedjoineduser.AvaterUrl = joineduser.AvaterUrl

			// スタンプの配列取得
			var stampNodes []model.Stamp
			if err := db.DB.Where("card_id = ?", card.Id).Find(&stampNodes).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					stampNodes = []model.Stamp{}
				} else {
					// return 500
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"message": "Database Error: " + err.Error(),
					})
				}
			}

			responseData = append(responseData, echo.Map{
				"title":         card.Title,
				"createdBy":     omitedcreateduser,
				"joinedUser":    omitedjoineduser,
				"createdAt":     card.CreatedAt,
				"updatedAt":     card.UpdatedAt,
				"currentDay":    card.CurrentDay,
				"isCompleted":   card.IsCompleted,
				"isDeleted":     card.IsDeleted,
				"stampNodes":    stampNodes,
				"backgroundUrl": card.BackgroundUrl,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"cards": responseData,
		})
	}
}

func CardCreate(c echo.Context) error {
	type Body struct {
		Title         string `json:"title"`
		JoinedUser    string `json:"JoinedUser"`
		StartDate     string `json:"startDate"`
		EndDate       string `json:"endDate"`
		CurrentDay    int    `json:"CurrentDay"`
		IsStampy      bool   `json:"IsStampy"`
		IsCompleted   bool   `json:"IsCompleted"`
		IsDeleted     bool   `json:"IsDeleted"`
		BackgroundUrl string `json:"BackgroundUrl"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	// create todo, return 201
	new := model.Stampcard{
		Title:         obj.Title,
		CreatedBy:     userid,
		JoinedUser:    obj.JoinedUser,
		StartDate:     obj.StartDate,
		EndDate:       obj.EndDate,
		CurrentDay:    obj.CurrentDay,
		IsStampy:      obj.IsStampy,
		IsCompleted:   obj.IsCompleted,
		IsDeleted:     obj.IsDeleted,
		BackgroundUrl: obj.BackgroundUrl,
	}
	db.DB.Create(&new)

	return c.JSON(http.StatusCreated, echo.Map{
		"id":            new.Id,
		"title":         new.Title,
		"CreatedBy":     new.CreatedBy,
		"JoinedUser":    new.JoinedUser,
		"startDate":     new.StartDate,
		"endDate":       new.EndDate,
		"CurrentDay":    new.CurrentDay,
		"IsStampy":      new.IsStampy,
		"IsCompleted":   new.IsCompleted,
		"IsDeleted":     new.IsDeleted,
		"BackgroundUrl": new.BackgroundUrl,
	})

}

func CardUpdate(c echo.Context) error {
	type Body struct {
		Title         string `json:"title"`
		CurrentDay    int    `json:"CurrentDay"`
		IsCompleted   bool   `json:"IsCompleted"`
		BackgroundUrl string `json:"BackgroundUrl"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	cardid := c.Param("id")
	var card model.Stampcard
	if err := db.DB.Where("id = ?", cardid).First(&card).Error; err != nil {
		// return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})

	} else {
		if card.CreatedBy == userid {
			obj := new(Body)
			if err := c.Bind(obj); err != nil {
				// return 400
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "Json Format Error: " + err.Error(),
				})
			}
			// update todo, return 204
			card.Title = obj.Title
			card.CurrentDay = obj.CurrentDay
			card.IsCompleted = obj.IsCompleted
			card.BackgroundUrl = obj.BackgroundUrl
			db.DB.Save(&user)
			return c.JSON(http.StatusCreated, echo.Map{
				"id":            card.Id,
				"title":         card.Title,
				"currentDay":    card.CurrentDay,
				"isCompleted":   card.IsCompleted,
				"backgroundUrl": card.BackgroundUrl,
			})

		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "No authority",
			})
		}
	}
}

func StampCreate(c echo.Context) error {
	type Body struct {
		StampImg string `json:"stamp"`
		Message  string `json:"title"`
		Nthday   int    `json:"nthday"`
		CardId   uint   `json:"cardId"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	new := model.Stamp{
		StampImg:  obj.StampImg,
		Message:   obj.Message,
		Nthday:    obj.Nthday,
		StampedBy: userid,
		CardId:    obj.CardId,
	}
	db.DB.Create(&new)

	return c.JSON(http.StatusCreated, echo.Map{
		"id":        new.Id,
		"stamp":     new.StampImg,
		"message":   new.Message,
		"nthday":    new.Nthday,
		"stampedBy": new.StampedBy,
		"cardId":    new.CardId,
		"createdAt": new.CreatedAt,
	})

}
