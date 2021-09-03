package handler

import (
	"encoding/csv"
	"go-api-report2/model"
	"io"
	"net/http"

	"github.com/jszwec/csvutil"
	"github.com/labstack/echo/v4"
)

func (h Handler) UploadUserProfile(c echo.Context) error {
	var result struct {
		Total int `json:"total,omitempty"`
	}

	slug := c.Param("slug")
	file, err := c.FormFile("file")
	user := ""
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	csvReader := csv.NewReader(src)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = '\t'

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return err
	}

	dec.Register(model.UnmarshalTime)

	var count int
	var data []interface{}
	filename := file.Filename

	switch slug {
	case "userprofileamlocsv":
		for {
			var w model.UserProfileAmlo
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}
	}
	count = len(data)
	err = h.reportService.PostDataUpload(data, count, slug, filename, user)
	if err != nil {
		h.logger.Errorf("[Handler] Report %s insert data error", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500107,
			Message: "Report insert data error",
		})
	}
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
