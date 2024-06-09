package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	ResponseCode string      `json:"ResponseCode"`
	ResponseDesc string      `json:"ResponseDesc"`
	ResponseData interface{} `json:"ResponseData"`
}

// getHTTPStatusFromErrorMessage extracts the HTTP status code from an error message.
// It splits the error message by "->" and parses the first part as an integer.
// If the parsing is successful, it returns the HTTP status code.
// If there is an error during parsing, it returns 0 and the parse error.
func getHTTPStatusFromErrorMessage(err error) (int, error) {
	// Split the error message by "->"
	splitErr := strings.Split(err.Error(), "->")

	// Parse the first part as an integer
	httpStatus, parseErr := strconv.Atoi(splitErr[0])
	if parseErr != nil {
		return 0, parseErr
	}

	return httpStatus, nil
}

// ErrorHandler handles errors and sends an appropriate response to the client.
// It takes an error and an echo.Context as input parameters.
// If the context's response has already been committed, the function returns without doing anything.
// Otherwise, it constructs a Response object with a response code of "0004".
// If the error is an echo.HTTPError, it extracts the error code and message from it.
// If the error is not an echo.HTTPError, it splits the error message and extracts the response code, description, and additional error details.
// Finally, it sends the constructed response with the appropriate HTTP status code.
func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	resp := Response{
		ResponseCode: "0004",
	}

	errCode := http.StatusBadRequest
	res, ok := err.(*echo.HTTPError)
	if ok {
		errCode = res.Code
		resp.ResponseDesc = fmt.Sprint(res.Message)
		resp.ResponseData = map[string]string{}
	}

	// fmt.Println("ERROR")
	// fmt.Println(err)

	if !ok {
		splitErr := strings.Split(err.Error(), "->")

		errCode, err = getHTTPStatusFromErrorMessage(err)
		if err != nil {
			errCode = http.StatusBadRequest
		}

		resp.ResponseCode = splitErr[1]
		resp.ResponseDesc = splitErr[3]
		resp.ResponseData = map[string]string{}
	}

	c.JSON(errCode, resp)
}

// ErrorMessage constructs an error message with the given error status code and error message.

// 01xx: Authentication errors
// 02xx: Authorization errors
// 03xx: Validation errors
// 04xx: Data errors
// 05xx: Third-party service errors
// 06xx: Internal errors
// 07xx: Generic errors
// 08xx: Error Proccess CRUD

func ErrorMessage(errorStatusCode string, err error) error {
	switch errorStatusCode {
	case "0001":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Invalid signing form")
	case "0002":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Client Id and Client Secret not found")
	case "0003":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Please try to create token later")
	case "0004":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Unauthorized")
	case "0005":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Data Already exist")
	case "0006":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "Data is not found")
	case "0009":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Generic service error")
	case "0011":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Data Invoice is not found")
	case "0012":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Error Validation")
	case "0013":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Wrong format data")
	case "0014":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Wrong Length data")
	case "0015":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Invoice is already paid")
	case "0016":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusServiceUnavailable, errorStatusCode, err, "The third-party service is currently experiencing issues. Please try again later.")
	case "0017":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusServiceUnavailable, errorStatusCode, err, "The third-party service does not provide data.")
	case "0018":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusServiceUnavailable, errorStatusCode, err, "GeneralError")
	case "0019":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusServiceUnavailable, errorStatusCode, err, "IDOrder tidak boleh kosong")
	case "0020":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusServiceUnavailable, errorStatusCode, err, "IDOrder tidak sesuai format")
	case "0101":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Invalid Signature")
	case "0102":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Invalid Timestamp")
	case "0103":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Invalid Token")
	case "0104":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Invalid API Key")
	case "0105":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Wrong Credentials")
	case "0106":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, errorStatusCode, err, "Inactive Account")
	case "0300":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Request body tidak valid")
	case "0301":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "NoHp harus diisi")
	case "0302":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "NoHp harus berupa angka")
	case "0303":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Panjang karakter NoHp tidak valid")
	case "0304":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "tanggal kuota tidak boleh kosong")
	case "0305":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "error convert int to string")
	case "0306":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "tanggal tidak sesuai format")
	case "0307":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "error time parse")
	case "0308":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "tidak dapat melakukan order, waktu sudah lebih dari jam cut off")
	case "0309":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "tanggal tidak sesuai order (7 hari)")
	case "0310":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "harga tabung belum di update")
	case "0311":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "kuota tidak cukup")
	case "0312":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "quantity order harus lebih besar dari 0")
	case "0313":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "status bayar harus berupa angka")
	case "0314":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "status bayar harus memiliki 2 digit")
	case "0400":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "Data pangkalan tidak ditemukan")
	case "0401":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "Data agen setting tidak ditemukan")
	case "0402":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "Data agen tidak ditemukan")
	case "0403":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "Data global param tidak ditemukan")
	case "0404":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "Data setting tidak ditemukan")
	case "0405":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "data range day tidak ditemukan")
	case "0406":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusNotFound, errorStatusCode, err, "data order tidak ditemukan")
	case "0500":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "koneksi ke pertamina terputus")
	case "0700":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "Generic errors")
	case "0800":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "gagal update lastNumber")
	case "0801":
		return fmt.Errorf("%d->%s->%v->%s", http.StatusBadRequest, errorStatusCode, err, "gagal membuat order")
	default:
		return fmt.Errorf("%d->%s->%v->%s", http.StatusUnauthorized, "0004", errors.New("Unauthorized"), "Unauthorized")
	}
}
