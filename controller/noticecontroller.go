package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nnnnn81/stampy-be/db"
	"github.com/nnnnn81/stampy-be/model"
	"gorm.io/gorm"
)

// お知らせ一覧取得(notice)

func NoticesShow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	var notices []model.Notice

	if err := db.DB.Where("receiver = ? and type = ?", userid, "notification").Find(&notices).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"notice": []model.Notice{},
			})
		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		var responseData []echo.Map
		for _, notice := range notices {
			// createdUser取得
			var sender model.User
			if err := db.DB.Where("id = ?", notice.Sender).First(&sender).Error; err != nil {
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
			if err := db.DB.Where("id = ?", notice.Receiver).First(&receiver).Error; err != nil {
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
				"id":         notice.Id,
				"type":       notice.Type,
				"title":      notice.Title,
				"stamp":      notice.Stamp,
				"message":    notice.Message,
				"nthDay":     notice.NthDay,
				"isLastDay":  notice.IsLastDay,
				"hrefPrefix": notice.HrefPrefix,
				"sender":     omitedsender,
				"receiver":   omitedreceiver,
				"read":       notice.Read,
				"createdAt":  notice.CreatedAt,
				"listType":   notice.ListType,
				"cardId":     notice.CardId,
				"letterId":   notice.LetterId,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"notice": responseData,
		})
	}
}

// お知らせ取得(id指定)
func NoticeShow(c echo.Context) error {
	noticeid := c.Param("id")
	var notice model.Notice

	if err := db.DB.Where("id = ? and type = ?", noticeid, "notification").Find(&notice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "notice not found",
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
		if err := db.DB.Where("id = ?", notice.Sender).First(&sender).Error; err != nil {
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
		if err := db.DB.Where("id = ?", notice.Receiver).First(&receiver).Error; err != nil {
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
			"id":         notice.Id,
			"type":       notice.Type,
			"title":      notice.Title,
			"stamp":      notice.Stamp,
			"message":    notice.Message,
			"nthDay":     notice.NthDay,
			"isLastDay":  notice.IsLastDay,
			"hrefPrefix": notice.HrefPrefix,
			"sender":     omitedsender,
			"receiver":   omitedreceiver,
			"read":       notice.Read,
			"createdAt":  notice.CreatedAt,
			"listType":   notice.ListType,
			"cardId":     notice.CardId,
			"letterId":   notice.LetterId,
		}

		return c.JSON(http.StatusOK, echo.Map{
			"notice": responseData,
		})
	}
}

// 通知作成(要求系)

func NoticeCreate(c echo.Context) error {
	type Body struct {
		StampId uint `json:"StampId"`
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

	var stamp model.Stamp
	if err := db.DB.Where("id = ?", obj.StampId).First(&stamp).Error; err != nil {
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
		var card model.Stampcard
		if err := db.DB.Where("id = ? and created_by = ?", stamp.CardId, userid).First(&card).Error; err != nil {
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
			if stamp.NthDay != card.Days {
				if card.IsStampy {
					// stampyの時、すぐにスタンプと受け取り通知作成
					// 一旦固定メッセージ
					stamp.StampImg = "🌟"
					stamp.Message = "えらい！"
					stamp.Stamped = true

					db.DB.Save(&stamp)

					newNotice := model.Notice{
						Type:       "notification",
						Title:      "スタンプが届いています",
						Stamp:      stamp.StampImg,
						Message:    stamp.Message,
						NthDay:     stamp.NthDay,
						HrefPrefix: "HrefPrefix",
						Sender:     card.JoinedUser,
						Receiver:   card.CreatedBy,
						ListType:   "receiver-dialog",
						CardId:     card.Id,
					}
					db.DB.Create(&newNotice)
					return c.JSON(http.StatusCreated, echo.Map{
						"notice": newNotice,
					})

				} else {
					newNotice := model.Notice{
						Type:       "notification",
						Title:      "スタンプを要求されています",
						HrefPrefix: "HrefPrefix",
						NthDay:     stamp.NthDay,
						Sender:     card.CreatedBy,
						Receiver:   card.JoinedUser,
						ListType:   "sender-dialog",
						CardId:     stamp.CardId,
					}
					db.DB.Create(&newNotice)

					return c.JSON(http.StatusCreated, echo.Map{
						"notices": newNotice,
					})
				}
			} else {
				if card.IsStampy {
					// stampyの時、すぐにレターと受け取り通知作成
					// 一旦固定メッセージ

					new := model.Letter{
						Type:       "letter",
						Title:      card.Title + "の完走レター",
						Stamp:      "🌟",
						Message:    "完走してえらい！",
						HrefPrefix: "/letter",
						Sender:     card.JoinedUser,
						Receiver:   card.CreatedBy,
						ListType:   "link",
						CardId:     card.Id,
					}
					db.DB.Create(&new)

					if !card.IsCompleted {
						card.IsCompleted = true
						card.LetterId = new.Id
					} else {
						return c.JSON(http.StatusBadRequest, echo.Map{
							"message": "this card is already finished",
						})
					}
					db.DB.Save(&card)

					stamp.StampImg = "🌟"
					stamp.Message = "完走！"
					stamp.Stamped = true

					db.DB.Save(&stamp)
					newNotice := model.Notice{
						Type:       "notification",
						Title:      card.Title + "の完走レターが届いています",
						Stamp:      "🌟",
						Message:    "完走してえらい！",
						NthDay:     stamp.NthDay,
						IsLastDay:  true,
						HrefPrefix: "HrefPrefix",
						Sender:     card.JoinedUser,
						Receiver:   card.CreatedBy,
						ListType:   "receiver-dialog",
						CardId:     card.Id,
						LetterId:   card.LetterId,
					}
					db.DB.Create(&newNotice)

					return c.JSON(http.StatusCreated, echo.Map{
						"notice": newNotice,
					})
				} else {
					newNotice := model.Notice{
						Type:       "notification",
						Title:      "レターを要求されています",
						HrefPrefix: "HrefPrefix",
						NthDay:     stamp.NthDay,
						IsLastDay:  true,
						Sender:     card.CreatedBy,
						Receiver:   card.JoinedUser,
						ListType:   "sender-dialog",
						CardId:     card.Id,
					}
					db.DB.Create(&newNotice)

					return c.JSON(http.StatusCreated, echo.Map{
						"notice": newNotice,
					})
				}
			}
		}
	}
}

// readの更新API
func NoticeReadUpdate(c echo.Context) error {
	notice_id := c.Param("id")

	var notice model.Notice
	if err := db.DB.Where("id = ?", notice_id).First(&notice).Error; err != nil {
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
		notice.Read = true
		db.DB.Save(&notice)
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "read",
		})
	}
}
