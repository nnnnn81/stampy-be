package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nnnnn81/stampy-be/db"
	"github.com/nnnnn81/stampy-be/model"
	"gorm.io/gorm"
)

func LettersShow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	var letters []model.Letter

	if err := db.DB.Where("receiver = ? and type = ?", userid, "letter").Find(&letters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"letter": []model.Letter{},
			})
		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		var responseData []echo.Map
		for _, letter := range letters {
			// createdUser取得
			var sender model.User
			if err := db.DB.Where("id = ?", letter.Sender).First(&sender).Error; err != nil {
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
			var receiver model.User
			if err := db.DB.Where("id = ?", letter.Receiver).First(&receiver).Error; err != nil {
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
			var omitedsender model.OmitUser
			omitedsender.Id = sender.Id
			omitedsender.Email = sender.Email
			omitedsender.Username = sender.Username
			omitedsender.AvatarUrl = sender.AvatarUrl
			var omitedreceiver model.OmitUser
			omitedreceiver.Id = receiver.Id
			omitedreceiver.Email = receiver.Email
			omitedreceiver.Username = receiver.Username
			omitedreceiver.AvatarUrl = receiver.AvatarUrl

			responseData = append(responseData, echo.Map{
				"id":         letter.Id,
				"type":       letter.Type,
				"title":      letter.Title,
				"stamp":      letter.Stamp,
				"message":    letter.Message,
				"hrefPrefix": letter.HrefPrefix,
				"sender":     omitedsender,
				"receiver":   omitedreceiver,
				"read":       letter.Read,
				"isVisible":  letter.IsVisible,
				"createdAt":  letter.CreatedAt,
				"listType":   letter.ListType,
				"cardId":     letter.CardId,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"letters": responseData,
		})
	}
}

// レター取得(id指定)

func LetterShow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	var letter model.Letter

	if err := db.DB.Where("receiver = ? and type = ? ", userid, "letter").Find(&letter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"letter": []model.Letter{},
			})
		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		// createdUser取得
		var sender model.User
		if err := db.DB.Where("id = ?", letter.Sender).First(&sender).Error; err != nil {
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
		var receiver model.User
		if err := db.DB.Where("id = ?", letter.Receiver).First(&receiver).Error; err != nil {
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
		var omitedsender model.OmitUser
		omitedsender.Id = sender.Id
		omitedsender.Email = sender.Email
		omitedsender.Username = sender.Username
		omitedsender.AvatarUrl = sender.AvatarUrl
		var omitedreceiver model.OmitUser
		omitedreceiver.Id = receiver.Id
		omitedreceiver.Email = receiver.Email
		omitedreceiver.Username = receiver.Username
		omitedreceiver.AvatarUrl = receiver.AvatarUrl

		responseData := echo.Map{
			"id":         letter.Id,
			"type":       letter.Type,
			"title":      letter.Title,
			"stamp":      letter.Stamp,
			"message":    letter.Message,
			"hrefPrefix": letter.HrefPrefix,
			"sender":     omitedsender,
			"receiver":   omitedreceiver,
			"read":       letter.Read,
			"createdAt":  letter.CreatedAt,
			"listType":   letter.ListType,
			"cardid":     letter.CardId,
		}

		return c.JSON(http.StatusOK, echo.Map{
			"letter": responseData,
		})
	}
}

// レター＆通知作成
func LetterCreate(c echo.Context) error {
	type Body struct {
		Message string
		Stamp   string
		CardId  uint
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
	if err := db.DB.Where("id = ? and joined_user = ?", obj.CardId, userid).First(&card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "card Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		var stamp model.Stamp
		if err := db.DB.Where("card_id = ? and nthday = ?", obj.CardId, card.Days).First(&stamp).Error; err != nil {
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

			newLetter := model.Letter{
				Type:       "letter",
				Title:      card.Title + "への完走レター",
				Stamp:      obj.Stamp,
				Message:    obj.Message,
				HrefPrefix: "/letter",
				Sender:     userid,
				Receiver:   card.CreatedBy,
				ListType:   "link",
				CardId:     obj.CardId,
			}
			db.DB.Create(&newLetter)

			if !card.IsCompleted {
				card.IsCompleted = true
				card.LetterId = newLetter.Id
			} else {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "this card is already finished",
				})
			}
			db.DB.Save(&card)

			stamp.StampImg = obj.Stamp
			stamp.Message = obj.Message
			stamp.Stamped = true

			db.DB.Save(&stamp)

			newNotice := model.Notice{
				Type:       "notification",
				Title:      card.Title + "への完走レターが届いています",
				Stamp:      obj.Stamp,
				Message:    obj.Message,
				CurrentDay: card.CurrentDay,
				IsLastDay:  true,
				HrefPrefix: "HrefPrefix",
				Sender:     userid,
				Receiver:   card.CreatedBy,
				ListType:   "receiver-dialog",
				CardId:     card.Id,
				LetterId:   card.LetterId,
			}
			db.DB.Create(&newNotice)

			return c.JSON(http.StatusCreated, echo.Map{
				"letter": newLetter,
			})
		}
	}
}

// readの更新API
func LetterReadUpdate(c echo.Context) error {
	letter_id := c.Param("id")

	var letter model.Letter
	if err := db.DB.Where("id = ?", letter_id).First(&letter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "letter Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		letter.Read = true
		db.DB.Save(&letter)
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "read",
		})
	}
}

// レターの開封処理
func VisibleUpdate(c echo.Context) error {
	letter_id := c.Param("id")

	var letter model.Letter
	if err := db.DB.Where("id = ?", letter_id).First(&letter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "letter Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		letter.IsVisible = true
		db.DB.Save(&letter)
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "letter opened",
		})
	}
}
