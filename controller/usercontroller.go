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
		AvaterUrl string `json:"avater_url"`
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
				AvaterUrl:      obj.AvaterUrl,
			}
			db.DB.Create(&new)

			// ペイロード作成
			claims := jwt.MapClaims{
				"id":  user.Id,
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
		omituser.AvaterUrl = user_.AvaterUrl
		// return 200
		return c.JSON(http.StatusCreated, omituser)

	}
}

func UserUpdate(c echo.Context) error {
	type Body struct {
		Username  string `json:"username"`
		AvaterUrl string `json:"avater_url"`
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
		user_.Username = obj.Username
		user_.AvaterUrl = obj.AvaterUrl
		db.DB.Save(&user)
		return c.JSON(http.StatusCreated, echo.Map{
			"id":         user_.Id,
			"username":   user_.Username,
			"email":      user_.Email,
			"avater_url": user_.AvaterUrl,
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
		hashedOldPass, err := util.HashPassword(obj.OldPass)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Password Hashing Error",
			})
		}
		if user_.HashedPassword == hashedOldPass {
			// update todo, return 204
			hashedNewPass, err := util.HashPassword(obj.NewPass)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Password Hashing Error",
				})
			}
			user_.HashedPassword = hashedNewPass
			db.DB.Save(&user)
			return c.JSON(http.StatusNoContent, echo.Map{})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "incorrect password",
			})
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
		return c.JSON(http.StatusNotFound, echo.Map{
			"doesUserExist": false,
		})

	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"doesUserExist": true,
		})
	}
}
