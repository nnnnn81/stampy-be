package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nnnnn81/stampy-be/db"
	"github.com/nnnnn81/stampy-be/model"
	"gorm.io/gorm"
)

func CardsShow(c echo.Context) error {
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
			if err := db.DB.Where("id = ?", card.JoinedUser).First(&joineduser).Error; err != nil {
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
			omitedcreateduser.AvatarUrl = createduser.AvatarUrl
			var omitedjoineduser model.OmitUser
			omitedjoineduser.Id = joineduser.Id
			omitedjoineduser.Email = joineduser.Email
			omitedjoineduser.Username = joineduser.Username
			omitedjoineduser.AvatarUrl = joineduser.AvatarUrl

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
				"id":            card.Id,
				"title":         card.Title,
				"createdBy":     omitedcreateduser,
				"joinedUser":    omitedjoineduser,
				"createdAt":     card.CreatedAt,
				"updatedAt":     card.UpdatedAt,
				"startDate":     card.StartDate,
				"endDate":       card.EndDate,
				"isCompleted":   card.IsCompleted,
				"isDeleted":     card.IsDeleted,
				"letterId":      card.LetterId,
				"stampNodes":    stampNodes,
				"backgroundUrl": card.BackgroundUrl,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"cards": responseData,
		})
	}
}
func CardShow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)
	cardid := c.Param("id")
	var card model.Stampcard
	if err := db.DB.Where("created_by = ? and id = ?", userid, cardid).First(&card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "card not found",
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		// createdUser取得
		log.Print(card)
		var createduser model.User
		if err := db.DB.Where("id = ?", card.CreatedBy).First(&createduser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// return 404
				return c.JSON(http.StatusNotFound, echo.Map{
					"message": "createUser Not Found",
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
		if err := db.DB.Where("id = ?", card.JoinedUser).First(&joineduser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// return 404
				return c.JSON(http.StatusNotFound, echo.Map{
					"message": "JoinedUser Not Found",
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
		omitedcreateduser.AvatarUrl = createduser.AvatarUrl
		var omitedjoineduser model.OmitUser
		omitedjoineduser.Id = joineduser.Id
		omitedjoineduser.Email = joineduser.Email
		omitedjoineduser.Username = joineduser.Username
		omitedjoineduser.AvatarUrl = joineduser.AvatarUrl

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
		responseData := echo.Map{
			"id":            card.Id,
			"title":         card.Title,
			"createdBy":     omitedcreateduser,
			"joinedUser":    omitedjoineduser,
			"createdAt":     card.CreatedAt,
			"updatedAt":     card.UpdatedAt,
			"startDate":     card.StartDate,
			"endDate":       card.EndDate,
			"isCompleted":   card.IsCompleted,
			"isDeleted":     card.IsDeleted,
			"letterId":      card.LetterId,
			"stampNodes":    stampNodes,
			"backgroundUrl": card.BackgroundUrl,
		}

		return c.JSON(http.StatusOK, responseData)
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
	joineduserEmail := obj.JoinedUser
	if obj.IsStampy {
		joineduserEmail = "stampy@gmail.com"
	}
	var joineduser model.User
	if err := db.DB.Where("email = ?", joineduserEmail).First(&joineduser).Error; err != nil {
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
	} else {
		// create card, return 201
		startDate, _ := time.Parse(time.RFC3339, obj.StartDate)
		endDate, _ := time.Parse(time.RFC3339, obj.EndDate)
		days := int(endDate.Sub(startDate).Hours() / 24)

		new := model.Stampcard{
			Title:         obj.Title,
			CreatedBy:     userid,
			JoinedUser:    joineduser.Id,
			StartDate:     obj.StartDate,
			EndDate:       obj.EndDate,
			Days:          days + 1,
			CurrentDay:    1,
			IsStampy:      obj.IsStampy,
			IsCompleted:   obj.IsCompleted,
			IsDeleted:     obj.IsDeleted,
			BackgroundUrl: obj.BackgroundUrl,
		}
		db.DB.Create(&new)

		for i := 0; i < days+1; i++ {
			newStamp := model.Stamp{
				StampImg:  "",
				Message:   "",
				NthDay:    i + 1,
				StampedBy: new.JoinedUser,
				CardId:    new.Id,
			}
			db.DB.Create(&newStamp)
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"id": new.Id,
		})
	}
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
			db.DB.Save(&card)
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
		Message  string `json:"message"`
		NthDay   int    `json:"nthday"`
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
	var card model.Stampcard
	if err := db.DB.Where("id = ?", obj.CardId).First(&card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Card Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		if card.JoinedUser != userid {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "you're not joined this card",
			})
		}
		var stamp model.Stamp
		if err := db.DB.Where("nth_day = ? and card_id = ?", obj.NthDay, obj.CardId).First(&stamp).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// return 404
				return c.JSON(http.StatusNotFound, echo.Map{
					"message": "stamp Not Found",
				})

			} else {
				// return 500
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Database Error: " + err.Error(),
				})
			}
		} else {
			var receiver model.User
			if err := db.DB.Where("id = ?", card.CreatedBy).First(&receiver).Error; err != nil {
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

			stamp.StampImg = obj.StampImg
			stamp.Message = obj.Message
			stamp.NthDay = obj.NthDay
			stamp.StampedBy = userid
			stamp.StampedTo = card.CreatedBy
			stamp.Stamped = true
			stamp.CardId = obj.CardId
			db.DB.Save(&stamp)
			newnotice := model.Notice{
				Type:       "notification",
				Title:      "スタンプが届いています",
				Stamp:      stamp.StampImg,
				Message:    stamp.Message,
				NthDay:     stamp.NthDay,
				HrefPrefix: "hrefPrefix",
				Sender:     userid,
				Receiver:   receiver.Id,
				ListType:   "receiver-dialog",
				CardId:     obj.CardId,
			}
			db.DB.Create(&newnotice)

			return c.JSON(http.StatusCreated, echo.Map{
				"id":        stamp.Id,
				"stamp":     stamp.StampImg,
				"message":   stamp.Message,
				"nthday":    stamp.NthDay,
				"stampedBy": stamp.StampedBy,
				"cardId":    stamp.CardId,
				"createdAt": stamp.CreatedAt,
			})

		}

	}
}

func CardDelete(c echo.Context) error {
	cardid := c.Param("id")

	var card model.Stampcard
	if err := db.DB.Where("id = ?", cardid).First(&card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Card Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		card.IsDeleted = true
		db.DB.Save(&card)
		return c.JSON(http.StatusNoContent, echo.Map{
			"message": "deleted",
		})
	}
}
