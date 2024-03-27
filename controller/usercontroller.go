package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nnnnn81/stampy-be/db"
	"github.com/nnnnn81/stampy-be/model"
	"github.com/nnnnn81/stampy-be/util"
	"gorm.io/gorm"
)

func Signup(c echo.Context) error {
	type Body struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		AvatarUrl string `json:"avatar_url"`
	}

	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	if util.HasEmptyField(obj, "Username", "Email", "Password") {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "missing request field",
		})
	}

	var user model.User
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPass, err := util.HashPassword(obj.Password)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Password Hashing Error",
				})
			}
			new := model.User{
				Username:       obj.Username,
				Email:          obj.Email,
				HashedPassword: hashedPass,
				AvatarUrl:      obj.AvatarUrl,
			}
			db.DB.Create(&new)

			// ペイロード作成
			claims := jwt.MapClaims{
				"id":  new.Id,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
			if err != nil {
				return err
			}
			// return 200
			return c.JSON(http.StatusOK, echo.Map{
				"token": tokenString,
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "email conflict",
		})
	}
}

func Login(c echo.Context) error {

	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// parse
	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	if util.HasEmptyField(obj, "Username", "Email", "Password") {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Missing Required Field",
		})
	}

	// ユーザーが存在するか
	var user model.User
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
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
		if err := util.ComparePasswords(user.HashedPassword, obj.Password); err != nil {
			// return 401
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid Password",
			})

		} else {
			// ペイロード作成
			claims := jwt.MapClaims{
				"id":  user.Id,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
			}
			// トークン生成
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// トークンに署名を付与
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
			if err != nil {
				return err
			}
			// return 200
			return c.JSON(http.StatusOK, echo.Map{
				"token": tokenString,
			})

		}
	}
}

func UserShow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)
	var user_ model.User
	if err := db.DB.Where("id = ?", userid).First(&user_).Error; err != nil {

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
		var omituser model.OmitUser
		omituser.Id = user_.Id
		omituser.Email = user_.Email
		omituser.Username = user_.Username
		omituser.AvatarUrl = user_.AvatarUrl
		// return 200
		return c.JSON(http.StatusCreated, omituser)

	}
}

func UserUpdate(c echo.Context) error {
	type Body struct {
		Username  string `json:"username"`
		AvatarUrl string `json:"avatarUrl"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	var user_ model.User
	if err := db.DB.Where("id = ?", userid).First(&user_).Error; err != nil {
		// return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})

	} else {

		// parse json
		obj := new(Body)
		if err := c.Bind(obj); err != nil {
			// return 400
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Json Format Error: " + err.Error(),
			})
		}
		// update todo, return 204
		if obj.Username != "" {
			user_.Username = obj.Username
		}
		if obj.AvatarUrl != "" {
			user_.AvatarUrl = obj.AvatarUrl
		}
		db.DB.Save(&user_)
		return c.JSON(http.StatusCreated, echo.Map{
			"id":        user_.Id,
			"username":  user_.Username,
			"email":     user_.Email,
			"avatarUrl": user_.AvatarUrl,
		})

	}
}

func UserPassUpdate(c echo.Context) error {
	type Body struct {
		OldPass string `json:"oldPass"`
		NewPass string `json:"newPass"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	var user_ model.User
	if err := db.DB.Where("id = ?", userid).First(&user_).Error; err != nil {
		// return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})

	} else {

		// parse json
		obj := new(Body)
		if err := c.Bind(obj); err != nil {
			// return 400
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Json Format Error: " + err.Error(),
			})
		}
		// パスワードの検証、ハッシュ化
		if err := util.ComparePasswords(user_.HashedPassword, obj.OldPass); err != nil {
			// return 401
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid Password",
			})

		} else {
			// update todo, return 204
			hashedNewPass, err := util.HashPassword(obj.NewPass)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Password Hashing Error",
				})
			} else {
				user_.HashedPassword = hashedNewPass
				db.DB.Save(&user_)
				return c.JSON(http.StatusOK, echo.Map{
					"message": "password changed",
				})
			}
		}
	}
}
func UserEmailCheck(c echo.Context) error {
	type Body struct {
		Email string `json:"email"`
	}
	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	var user model.User
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
		// return 200
		return c.JSON(http.StatusOK, echo.Map{
			"doesUserExist": false,
		})

	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"doesUserExist": true,
		})
	}
}

func UserTotal(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	var receivedStamp, challengeCard, completedCard, receivedLetter int64

	db.DB.Model(&model.Stamp{}).Where("stamped_to = ? and stamped = ?", userid, true).Count(&receivedStamp)

	db.DB.Model(&model.Stampcard{}).Where("created_by = ?", userid).Count(&challengeCard)

	db.DB.Model(&model.Stampcard{}).Where("created_by = ? and is_completed = ?", userid, true).Count(&completedCard)

	db.DB.Model(&model.Letter{}).Where("receiver = ?", userid).Count(&receivedLetter)

	response := map[string]int64{
		"receivedStamp":  receivedStamp,
		"challengeCard":  challengeCard,
		"completedCard":  completedCard,
		"receivedLetter": receivedLetter,
	}

	return c.JSON(http.StatusOK, response)
}
