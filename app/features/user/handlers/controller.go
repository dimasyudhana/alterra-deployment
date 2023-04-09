package handlers

import (
	"log"
	"net/http"

	user "github.com/dimasyudhana/latihan-deployment.git/app/features/user"
	"github.com/dimasyudhana/latihan-deployment.git/helper"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service user.UseCase
}

func New(us user.UseCase) user.Handler {
	return &UserController{
		service: us,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			code, res := helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil)
			return c.JSON(code, res)
		}

		// Cek apakah username kurang dari 3 huruf, jika iya maka diulangi
		if len(input.Username) < 3 {
			code, res := helper.ResponseFormat(http.StatusBadRequest, "username minimal memiliki 3 karakter", nil)
			return c.JSON(code, res)
		}

		// Cek apakah nomor telepon sudah terdaftar
		users, err := uc.service.FindByPhone(input.Phone)
		if err != nil {
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "gagal mengambil data user", nil)
			return c.JSON(code, res)
		}
		if len(users) > 0 {
			code, res := helper.ResponseFormat(http.StatusConflict, "nomor telepon telah terdaftar", nil)
			return c.JSON(code, res)
		}

		// isPasswordValid apakah password mengandung unique char dan atleast 8 char.
		if !helper.IsPasswordValid(input.Password) {
			code, res := helper.ResponseFormat(http.StatusBadRequest, "password harus terdiri dari huruf besar (A-Z), huruf kecil (a-z), nomor (0-9), simbol istimewa dan minimal memiliki 8 karakter", nil)
			return c.JSON(code, res)
		}

		// Cek apakah password dan confirm_password sama
		if !helper.IsPasswordMatched(input.Password, input.Confirm_password) {
			code, res := helper.ResponseFormat(http.StatusBadRequest, "password dan confirm password tidak sama", nil)
			return c.JSON(code, res)
		}

		err = uc.service.Register(user.Core{Username: input.Username, Phone: input.Phone, Email: input.Email, Password: input.Password, Confirm_password: input.Confirm_password})
		if err != nil {
			if err.Error() == "register logic error: phone number "+input.Phone+" has been registered" {
				code, res := helper.ResponseFormat(http.StatusConflict, err.Error(), nil)
				return c.JSON(code, res)
			}
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "gagal menambahkan data", nil)
			return c.JSON(code, res)
		}

		code, res := helper.ResponseFormat(http.StatusCreated, "sukses menambahkan data", nil)
		return c.JSON(code, res)
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		res, err := uc.service.Login(input.Phone, input.Password)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		var result = new(LoginResponse)
		token := helper.GenerateJWT(res.Phone)
		result.Nama = res.Username
		result.Phone = res.Phone
		result.Token = token

		return c.JSON(helper.ResponseFormat(http.StatusOK, "sukses login, gunakan token ini pada akses API selanjutnya.", result))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		update := new(UpdateResponse)
		if err := c.Bind(update); err != nil {
			log.Println("Failed to bind request body", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request body",
			})
		}

		updatedUser, err := uc.service.UpdateByPhone(update.Phone, update.Username, update.Email)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, err.Error(), nil))
		}

		response := map[string]interface{}{
			"username": updatedUser.Username,
			"phone":    updatedUser.Phone,
			"email":    updatedUser.Email,
		}

		return c.JSON(helper.ResponseFormat(http.StatusAccepted, "sukses melakukan update", response))
	}
}
